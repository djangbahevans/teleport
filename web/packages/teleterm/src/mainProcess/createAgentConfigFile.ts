/**
 * Copyright 2023 Gravitational, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { access, rm, writeFile } from 'node:fs/promises';
import path from 'node:path';

import * as constants from 'shared/constants';

import { RootClusterUri, routing } from 'teleterm/ui/uri';
import { RuntimeSettings } from 'teleterm/mainProcess/types';

export interface CreateAgentConfigFileArgs {
  rootClusterUri: RootClusterUri;
  proxy: string;
  token: string;
  username: string;
  fileServerPort: number;
}

export async function createAgentJoinedFile(
  runtimeSettings: RuntimeSettings,
  rootClusterUri: RootClusterUri
): Promise<void> {
  const { agentJoined } = generateAgentConfigPaths(
    runtimeSettings,
    rootClusterUri
  );

  await writeFile(agentJoined, '');
}

export async function removeAgentDirectory(
  runtimeSettings: RuntimeSettings,
  rootClusterUri: RootClusterUri
): Promise<void> {
  const { agentDirectory } = generateAgentConfigPaths(
    runtimeSettings,
    rootClusterUri
  );
  // `force` ignores exceptions if path does not exist
  await rm(agentDirectory, { recursive: true, force: true });
}

export async function isAgentJoinedFileCreated(
  runtimeSettings: RuntimeSettings,
  rootClusterUri: RootClusterUri
): Promise<boolean> {
  const { agentJoined } = generateAgentConfigPaths(
    runtimeSettings,
    rootClusterUri
  );
  try {
    await access(agentJoined);
    return true;
  } catch (e) {
    if (e.code === 'ENOENT') {
      return false;
    }
    throw e;
  }
}

/**
 * Returns agent config paths.
 * @param runtimeSettings must not come from the renderer process.
 * Otherwise, the generated paths may point outside the user's data directory.
 * @param rootClusterUri may be passed from the renderer process.
 */
export function generateAgentConfigPaths(
  runtimeSettings: RuntimeSettings,
  rootClusterUri: RootClusterUri
): {
  agentDirectory: string;
  agentJoined: string;
  logsDirectory: string;
  dataDirectory: string;
} {
  const parsed = routing.parseClusterUri(rootClusterUri);
  if (!parsed?.params?.rootClusterId) {
    throw new Error(`Incorrect root cluster URI: ${rootClusterUri}`);
  }

  const agentDirectory = getAgentDirectoryOrThrow(
    runtimeSettings.userDataDir,
    parsed.params.rootClusterId
  );
  const agentJoined = path.resolve(agentDirectory, 'agent_joined');
  const dataDirectory = path.resolve(agentDirectory, 'data');
  const logsDirectory = path.resolve(agentDirectory, 'logs');

  return {
    agentDirectory,
    agentJoined,
    dataDirectory,
    logsDirectory,
  };
}

export function getAgentsDir(userDataDir: string): string {
  // Why not put agentsDir into runtimeSettings? That's because we don't want the renderer to have
  // access to this value as it could lead to bad security practices.
  //
  // If agentsDir was sent from the renderer to tshd and the main process, those recipients could
  // not trust that agentsDir has not been tampered with. Instead, the renderer should merely send
  // the root cluster URI and the recipients should build the path to the specific agent dir from
  // that, with agentsDir being supplied out of band.
  return path.resolve(userDataDir, 'agents');
}

function getAgentDirectoryOrThrow(
  userDataDir: string,
  profileName: string
): string {
  const agentsDir = getAgentsDir(userDataDir);
  const resolved = path.resolve(agentsDir, profileName);

  // check if the path doesn't contain any unexpected segments
  const isValidPath =
    path.dirname(resolved) === agentsDir &&
    path.basename(resolved) === profileName;
  if (!isValidPath) {
    throw new Error(`The agent config path is incorrect: ${resolved}`);
  }
  return resolved;
}

type AgentConfig = object;

export function generateAgentConfig(
  runtimeSettings: RuntimeSettings,
  args: CreateAgentConfigFileArgs
): AgentConfig {
  const { dataDirectory } = generateAgentConfigPaths(
    runtimeSettings,
    args.rootClusterUri
  );
  const labels = {
    [constants.ConnectMyComputerNodeOwnerLabel]: args.username,
  };

  return generateConfig({
    proxy: args.proxy,
    token: args.token,
    nodename: runtimeSettings.hostname,
    dataDir: dataDirectory,
    labels,
    fileServerPort: args.fileServerPort,
  });
}

export function generateConfig(args: {
  proxy: string;
  token: string;
  nodename: string;
  dataDir: string;
  labels: Record<string, string>;
  fileServerPort: number;
}): AgentConfig {
  return {
    version: 'v3',
    teleport: {
      nodename: args.nodename,
      data_dir: args.dataDir,
      join_params: {
        token_name: args.token,
        method: 'token',
      },
      proxy_server: args.proxy,
      log: {
        output: 'stderr',
        severity: 'INFO',
        format: {
          output: 'text',
        },
      },
    },
    auth_service: {
      enabled: 'no',
    },
    ssh_service: {
      enabled: 'yes',
      labels: args.labels,
    },
    proxy_service: {
      enabled: 'no',
    },
    app_service: {
      enabled: 'yes',
      apps: [
        {
          //TODO: make it random
          name: 'file-sharing',
          uri: `https://127.0.0.1:${args.fileServerPort}`,
          insecure_skip_verify: true,
          labels: {
            env: 'test',
          },
        },
      ],
    },
  };
}
