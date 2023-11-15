/*
Copyright 2023 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import path from 'node:path';
import childProcess, { ChildProcess } from 'node:child_process';
import fs from 'node:fs/promises';
import { tmpdir } from 'node:os';

import Logger, { NullService } from 'teleterm/logger';

import { CreateAgentConfigFileArgs } from 'teleterm/mainProcess/createAgentConfigFile';

import * as mocks from '../fixtures/mocks';
import { AgentProcessState, RuntimeSettings } from '../types';

import { AgentRunner } from './agentRunner';

const makeRuntimeSettings = (settings?: Partial<RuntimeSettings>) =>
  mocks.makeRuntimeSettings({
    agentBinaryPath,
    userDataDir,
    logsDir,
    ...settings,
  });

beforeAll(async () => {
  // Create a temp dir for user data dir. The cleanup daemon is going to store logs there.
  userDataDir = await fs.mkdtemp(
    path.join(tmpdir(), 'agent-cleanup-daemon-test-logs')
  );
  logsDir = path.join(userDataDir, 'logs');
});

afterAll(async () => {
  await fs.rm(userDataDir, { recursive: true, force: true });
});

beforeEach(() => {
  Logger.init(new NullService());
  jest.spyOn(childProcess, 'fork');
});

afterEach(() => {
  jest.restoreAllMocks();
});

let userDataDir: string;
let logsDir: string;
const agentBinaryPath = path.join(__dirname, 'agentTestProcess.mjs');
const agentCleanupDaemonPath = path.join(
  __dirname,
  '..',
  '..',
  'agentCleanupDaemon',
  'agentCleanupDaemon.js'
);
const clusterArgs: CreateAgentConfigFileArgs = {
  rootClusterUri: '/clusters/cluster.local',
  proxy: 'cluster.local:3080',
  token: '',
  username: 'adrian',
  fileServerPort: 6789,
};

test('agent process and cleanup daemon start with correct arguments', async () => {
  const agentRunner = new AgentRunner(
    makeRuntimeSettings(),
    agentCleanupDaemonPath,
    () => {}
  );

  try {
    const agentProcess = await agentRunner.start(clusterArgs);
    await new Promise(resolve => agentProcess.once('spawn', resolve));

    expect(childProcess.fork).toHaveBeenCalled();
    const cleanupDaemon = (
      childProcess.fork as jest.MockedFunction<typeof childProcess.fork>
    ).mock.results[0].value;

    //TODO: uncomment
    // expect(agentProcess.spawnargs).toEqual([
    //   agentBinaryPath,
    //   'start',
    //   '--config-string=eyJ2ZXJzaW9uIjoidjMiLCJ0ZWxlcG9ydCI6eyJub2RlbmFtZSI6InN0YWdpbmctbWFjLW1pbmkiLCJkYXRhX2RpciI6Ii92YXIvZm9sZGVycy95NS95cWc4eHo1NTVfdjd4ZnNyMHduNmI0cTgwMDAwZ24vVC9hZ2VudC1jbGVhbnVwLWRhZW1vbi10ZXN0LWxvZ3MwbVU5U2QvYWdlbnRzL2NsdXN0ZXIubG9jYWwvZGF0YSIsImpvaW5fcGFyYW1zIjp7InRva2VuX25hbWUiOiIiLCJtZXRob2QiOiJ0b2tlbiJ9LCJwcm94eV9zZXJ2ZXIiOiJjbHVzdGVyLmxvY2FsOjMwODAiLCJsb2ciOnsib3V0cHV0Ijoic3RkZXJyIiwic2V2ZXJpdHkiOiJJTkZPIiwiZm9ybWF0Ijp7Im91dHB1dCI6InRleHQifX19LCJhdXRoX3NlcnZpY2UiOnsiZW5hYmxlZCI6Im5vIn0sInNzaF9zZXJ2aWNlIjp7ImVuYWJsZWQiOiJ5ZXMiLCJsYWJlbHMiOnsidGVsZXBvcnQuZGV2L2Nvbm5lY3QtbXktY29tcHV0ZXIvb3duZXIiOiJhZHJpYW4ifX0sInByb3h5X3NlcnZpY2UiOnsiZW5hYmxlZCI6Im5vIn0sImFwcF9zZXJ2aWNlIjp7ImVuYWJsZWQiOiJ5ZXMiLCJhcHBzIjpbeyJuYW1lIjoiZmlsZS1zaGFyaW5nIiwidXJpIjoiaHR0cHM6Ly8xMjcuMC4wLjE6Njc4OSIsImluc2VjdXJlX3NraXBfdmVyaWZ5Ijp0cnVlLCJsYWJlbHMiOnsiZW52IjoidGVzdCJ9fV19fQ==',
    //   '--insecure',
    // ]);
    expect(cleanupDaemon.spawnargs).toEqual([
      process.argv[0], // path to Node.js bin
      agentCleanupDaemonPath,
      agentProcess.pid.toString(),
      process.pid.toString(),
      clusterArgs.rootClusterUri,
      logsDir,
    ]);
  } finally {
    await agentRunner.killAll();
  }
});

test('previous agent process is killed when a new one is started', async () => {
  const agentRunner = new AgentRunner(
    makeRuntimeSettings(),
    agentCleanupDaemonPath,
    () => {}
  );

  try {
    const firstProcess = await agentRunner.start(clusterArgs);
    await agentRunner.start(clusterArgs);

    expect(firstProcess.killed).toBeTruthy();
  } finally {
    await agentRunner.killAll();
  }
});

test('status updates are sent on a successful start', async () => {
  const updateSender = jest.fn();
  const agentRunner = new AgentRunner(
    makeRuntimeSettings(),
    agentCleanupDaemonPath,
    updateSender
  );

  try {
    expect(agentRunner.getState(clusterArgs.rootClusterUri)).toBeUndefined();
    const agentProcess = await agentRunner.start(clusterArgs);
    expect(agentRunner.getState(clusterArgs.rootClusterUri)).toStrictEqual({
      status: 'not-started',
    } as AgentProcessState);

    await new Promise(resolve => agentProcess.once('spawn', resolve));

    const runningState: AgentProcessState = { status: 'running' };
    expect(agentRunner.getState(clusterArgs.rootClusterUri)).toStrictEqual(
      runningState
    );
    expect(updateSender).toHaveBeenCalledWith(
      clusterArgs.rootClusterUri,
      runningState
    );

    await agentRunner.kill(clusterArgs.rootClusterUri);

    // Since the agent changes status on the close event and not the exit event, we must wait for
    // this to occur.
    await expect(
      () => agentRunner.getState(clusterArgs.rootClusterUri).status === 'exited'
    ).toEventuallyBeTrue({
      waitFor: 2000,
      tick: 10,
    });

    const exitedState: AgentProcessState = {
      status: 'exited',
      code: null,
      logs: undefined,
      exitedSuccessfully: true,
      signal: 'SIGTERM',
    };
    expect(agentRunner.getState(clusterArgs.rootClusterUri)).toStrictEqual(
      exitedState
    );
    expect(updateSender).toHaveBeenCalledWith(
      clusterArgs.rootClusterUri,
      exitedState
    );

    expect(updateSender).toHaveBeenCalledTimes(2);
  } finally {
    await agentRunner.killAll();
  }
});

test('status updates are sent on a failed start', async () => {
  const updateSender = jest.fn();
  const nonExisingPath = path.join(
    __dirname,
    'agentTestProcess-nonExisting.mjs'
  );
  const agentRunner = new AgentRunner(
    makeRuntimeSettings({
      agentBinaryPath: nonExisingPath,
    }),
    agentCleanupDaemonPath,
    updateSender
  );

  try {
    const agentProcess = await agentRunner.start(clusterArgs);
    await new Promise(resolve => agentProcess.on('error', resolve));

    expect(updateSender).toHaveBeenCalledTimes(1);
    const errorState: AgentProcessState = {
      status: 'error',
      message: expect.stringContaining('ENOENT'),
    };
    expect(agentRunner.getState(clusterArgs.rootClusterUri)).toStrictEqual(
      errorState
    );
    expect(updateSender).toHaveBeenCalledWith(
      clusterArgs.rootClusterUri,
      errorState
    );
  } finally {
    await agentRunner.killAll();
  }
});

test('cleanup daemon stops together with agent process', async () => {
  const agentRunner = new AgentRunner(
    makeRuntimeSettings(),
    agentCleanupDaemonPath,
    () => {}
  );

  try {
    const agent = await agentRunner.start(clusterArgs);
    await new Promise(resolve => agent.once('spawn', resolve));

    expect(childProcess.fork).toHaveBeenCalled();
    const cleanupDaemon = (
      childProcess.fork as jest.MockedFunction<typeof childProcess.fork>
    ).mock.results[0].value;

    await agentRunner.kill(clusterArgs.rootClusterUri);

    expect(isRunning(agent)).toBe(false);
    // The cleanup daemon is killed from within an event listener, so it won't be killed
    // immediately.
    await expect(() => !isRunning(cleanupDaemon)).toEventuallyBeTrue({
      waitFor: 2000,
      tick: 10,
    });
  } finally {
    await agentRunner.killAll();
  }
});

test('agent cleanup daemon is not spawned on failed agent start', async () => {
  const nonExisingPath = path.join(
    __dirname,
    'agentTestProcess-nonExisting.mjs'
  );
  const agentRunner = new AgentRunner(
    makeRuntimeSettings({
      agentBinaryPath: nonExisingPath,
    }),
    agentCleanupDaemonPath,
    () => {}
  );

  try {
    const agent = await agentRunner.start(clusterArgs);
    await new Promise(resolve => agent.on('error', resolve));

    expect(isRunning(agent)).toBe(false);
    expect(childProcess.fork).not.toHaveBeenCalled();
  } finally {
    await agentRunner.killAll();
  }
});

// It'd be nice to test a situation where the cleanup daemon fails to spawn, but it's unclear how to
// test it when using `fork` to spawn the cleanup daemon.
test('agent is killed if cleanup daemon exits', async () => {
  const agentRunner = new AgentRunner(
    makeRuntimeSettings(),
    agentCleanupDaemonPath,
    () => {}
  );

  try {
    const agentProcess = await agentRunner.start(clusterArgs);
    await new Promise(resolve => agentProcess.once('spawn', resolve));

    expect(childProcess.fork).toHaveBeenCalled();
    const cleanupDaemon: ChildProcess = (
      childProcess.fork as jest.MockedFunction<typeof childProcess.fork>
    ).mock.results[0].value;

    cleanupDaemon.kill('SIGKILL');

    await expect(() => !isRunning(agentProcess)).toEventuallyBeTrue({
      waitFor: 2000,
      tick: 10,
    });

    expect(childProcess.fork).toHaveBeenCalled();
  } finally {
    await agentRunner.killAll();
  }
});

const isRunning = (process: ChildProcess) =>
  process.exitCode === null && process.signalCode === null;
