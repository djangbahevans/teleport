/**
 * Teleport
 * Copyright (C) 2024 Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import { equalsDeep } from 'shared/utils/highbar';
import { Option } from 'shared/components/Select';

import {
  KubernetesResource,
  Labels,
  Role,
  RoleConditions,
} from 'teleport/services/resources';
import { Label as UILabel } from 'teleport/components/LabelsInput/LabelsInput';
import {
  CreateDBUserMode,
  CreateHostUserMode,
  KubernetesResourceKind,
  KubernetesVerb,
  RequireMFAType,
  ResourceKind,
  RoleOptions,
  Rule,
  Verb,
} from 'teleport/services/resources/types';

import { defaultOptions } from './withDefaults';

export type StandardEditorModel = {
  roleModel: RoleEditorModel;
  /**
   * Will be true if fields have been modified from the original.
   */
  isDirty: boolean;
};

/**
 * A temporary representation of the role, reflecting the structure of standard
 * editor UI. Since the standard editor UI structure doesn't directly represent
 * the structure of the role resource, we introduce this intermediate model.
 */
export type RoleEditorModel = {
  metadata: MetadataModel;
  accessSpecs: AccessSpec[];
  rules: RuleModel[];
  options: OptionsModel;
  /**
   * Indicates whether the current resource, as described by YAML, is
   * accurately represented by this editor model. If it's not, the user needs
   * to agree to reset it to a compatible resource before editing it in the
   * structured editor.
   */
  requiresReset: boolean;
};

export type MetadataModel = {
  name: string;
  description?: string;
  revision?: string;
};

/** A model for access specifications section. */
export type AccessSpec =
  | KubernetesAccessSpec
  | ServerAccessSpec
  | AppAccessSpec
  | DatabaseAccessSpec
  | WindowsDesktopAccessSpec;

/**
 * A base for all access specification section models. Contains a type
 * discriminator field.
 */
type AccessSpecBase<T extends AccessSpecKind> = {
  /**
   * Determines kind of resource that is accessed using this spec. Intended to
   * be mostly consistent with UnifiedResources.kind, but that has no real
   * meaning on the server side; we needed some discriminator, so we picked
   * this one.
   */
  kind: T;
};

export type AccessSpecKind =
  | 'node'
  | 'kube_cluster'
  | 'app'
  | 'db'
  | 'windows_desktop';

/** Model for the Kubernetes access specification section. */
export type KubernetesAccessSpec = AccessSpecBase<'kube_cluster'> & {
  groups: readonly Option[];
  labels: UILabel[];
  resources: KubernetesResourceModel[];
};

export type KubernetesResourceModel = {
  /** Autogenerated ID to be used with the `key` property. */
  id: string;
  kind: KubernetesResourceKindOption;
  name: string;
  namespace: string;
  verbs: readonly KubernetesVerbOption[];
};

type KubernetesResourceKindOption = Option<KubernetesResourceKind, string>;

/**
 * All possible resource kind drop-down options. This array needs to be kept in
 * sync with `KubernetesResourcesKinds` in `api/types/constants.go.
 */
export const kubernetesResourceKindOptions: KubernetesResourceKindOption[] = [
  // The "any kind" option goes first.
  { value: '*', label: 'Any kind' },

  // The rest is sorted by label.
  ...(
    [
      { value: 'pod', label: 'Pod' },
      { value: 'secret', label: 'Secret' },
      { value: 'configmap', label: 'ConfigMap' },
      { value: 'namespace', label: 'Namespace' },
      { value: 'service', label: 'Service' },
      { value: 'serviceaccount', label: 'ServiceAccount' },
      { value: 'kube_node', label: 'Node' },
      { value: 'persistentvolume', label: 'PersistentVolume' },
      { value: 'persistentvolumeclaim', label: 'PersistentVolumeClaim' },
      { value: 'deployment', label: 'Deployment' },
      { value: 'replicaset', label: 'ReplicaSet' },
      { value: 'statefulset', label: 'Statefulset' },
      { value: 'daemonset', label: 'DaemonSet' },
      { value: 'clusterrole', label: 'ClusterRole' },
      { value: 'kube_role', label: 'Role' },
      { value: 'clusterrolebinding', label: 'ClusterRoleBinding' },
      { value: 'rolebinding', label: 'RoleBinding' },
      { value: 'cronjob', label: 'Cronjob' },
      { value: 'job', label: 'Job' },
      {
        value: 'certificatesigningrequest',
        label: 'CertificateSigningRequest',
      },
      { value: 'ingress', label: 'Ingress' },
    ] as const
  ).toSorted((a, b) => a.label.localeCompare(b.label)),
];

const optionsToMap = <K, V>(opts: Option<K, V>[]) =>
  new Map(opts.map(o => [o.value, o]));

const kubernetesResourceKindOptionsMap = optionsToMap(
  kubernetesResourceKindOptions
);

type KubernetesVerbOption = Option<KubernetesVerb, string>;
/**
 * All possible Kubernetes verb drop-down options. This array needs to be kept
 * in sync with `KubernetesVerbs` in `api/types/constants.go.
 */
export const kubernetesVerbOptions: KubernetesVerbOption[] = [
  // The "any kind" option goes first.
  { value: '*', label: 'All verbs' },

  // The rest is sorted.
  ...(
    [
      'get',
      'create',
      'update',
      'patch',
      'delete',
      'list',
      'watch',
      'deletecollection',

      // TODO(bl-nero): These are actually not k8s verbs, but they are allowed
      // in our config. We may want to explain them in the UI somehow.
      'exec',
      'portforward',
    ] as const
  )
    .toSorted((a, b) => a.localeCompare(b))
    .map(stringToOption),
];
const kubernetesVerbOptionsMap = optionsToMap(kubernetesVerbOptions);

type ResourceKindOption = Option<ResourceKind, string>;
export const resourceKindOptions: ResourceKindOption[] = Object.values(
  ResourceKind
)
  .toSorted()
  .map(stringToOption);
const resourceKindOptionsMap = optionsToMap(resourceKindOptions);

type VerbOption = Option<Verb, string>;
export const verbOptions: VerbOption[] = (
  [
    '*',
    'create',
    'create_enroll_token',
    'delete',
    'enroll',
    'list',
    'read',
    'readnosecrets',
    'rotate',
    'update',
    'use',
  ] as const
).map(stringToOption);
const verbOptionsMap = optionsToMap(verbOptions);

/** Model for the server access specification section. */
export type ServerAccessSpec = AccessSpecBase<'node'> & {
  labels: UILabel[];
  logins: readonly Option[];
};

export type AppAccessSpec = AccessSpecBase<'app'> & {
  labels: UILabel[];
  awsRoleARNs: string[];
  azureIdentities: string[];
  gcpServiceAccounts: string[];
};

export type DatabaseAccessSpec = AccessSpecBase<'db'> & {
  labels: UILabel[];
  names: readonly Option[];
  users: readonly Option[];
  roles: readonly Option[];
};

export type WindowsDesktopAccessSpec = AccessSpecBase<'windows_desktop'> & {
  labels: UILabel[];
  logins: readonly Option[];
};

export type RuleModel = {
  /** Autogenerated ID to be used with the `key` property. */
  id: string;
  resources: readonly ResourceKindOption[];
  verbs: readonly VerbOption[];
};

export type OptionsModel = {
  maxSessionTTL: string;
  clientIdleTimeout: string;
  disconnectExpiredCert: boolean;
  requireMFAType: RequireMFATypeOption;
  createHostUserMode: CreateHostUserModeOption;
  createDBUser: boolean;
  createDBUserMode: CreateDBUserModeOption;
  desktopClipboard: boolean;
  createDesktopUser: boolean;
  desktopDirectorySharing: boolean;
};

type RequireMFATypeOption = Option<RequireMFAType>;
export const requireMFATypeOptions: RequireMFATypeOption[] = [
  { value: false, label: 'No' },
  { value: true, label: 'Yes' },
  { value: 'hardware_key', label: 'Hardware Key' },
  { value: 'hardware_key_touch', label: 'Hardware Key (touch)' },
  {
    value: 'hardware_key_touch_and_pin',
    label: 'Hardware Key (touch and PIN)',
  },
];
const requireMFATypeOptionsMap = optionsToMap(requireMFATypeOptions);

type CreateHostUserModeOption = Option<CreateHostUserMode>;
export const createHostUserModeOptions: CreateHostUserModeOption[] = [
  { value: '', label: 'Unspecified' },
  { value: 'off', label: 'Off' },
  { value: 'keep', label: 'Keep' },
  { value: 'insecure-drop', label: 'Drop (insecure)' },
];
const createHostUserModeOptionsMap = optionsToMap(createHostUserModeOptions);

type CreateDBUserModeOption = Option<CreateDBUserMode>;
export const createDBUserModeOptions: CreateDBUserModeOption[] = [
  { value: '', label: 'Unspecified' },
  { value: 'off', label: 'Off' },
  { value: 'keep', label: 'Keep' },
  { value: 'best_effort_drop', label: 'Drop (best effort)' },
];
const createDBUserModeOptionsMap = optionsToMap(createDBUserModeOptions);

const roleVersion = 'v7';

/**
 * Returns the role object with required fields defined with empty values.
 */
export function newRole(): Role {
  return {
    kind: 'role',
    metadata: {
      name: 'new_role_name',
    },
    spec: {
      allow: {},
      deny: {},
      options: defaultOptions(),
    },
    version: roleVersion,
  };
}

export function newAccessSpec(kind: 'node'): ServerAccessSpec;
export function newAccessSpec(kind: 'kube_cluster'): KubernetesAccessSpec;
export function newAccessSpec(kind: 'app'): AppAccessSpec;
export function newAccessSpec(kind: 'db'): DatabaseAccessSpec;
export function newAccessSpec(
  kind: 'windows_desktop'
): WindowsDesktopAccessSpec;
export function newAccessSpec(kind: AccessSpecKind): AppAccessSpec;
export function newAccessSpec(kind: AccessSpecKind): AccessSpec {
  switch (kind) {
    case 'node':
      return { kind: 'node', labels: [], logins: [] };
    case 'kube_cluster':
      return { kind: 'kube_cluster', groups: [], labels: [], resources: [] };
    case 'app':
      return {
        kind: 'app',
        labels: [],
        awsRoleARNs: [],
        azureIdentities: [],
        gcpServiceAccounts: [],
      };
    case 'db':
      return { kind: 'db', labels: [], names: [], users: [], roles: [] };
    case 'windows_desktop':
      return { kind: 'windows_desktop', labels: [], logins: [] };
    default:
      kind satisfies never;
  }
}

export function newKubernetesResourceModel(): KubernetesResourceModel {
  return {
    id: crypto.randomUUID(),
    kind: kubernetesResourceKindOptions.find(k => k.value === '*'),
    name: '*',
    namespace: '*',
    verbs: [],
  };
}

export function newRuleModel(): RuleModel {
  return {
    id: crypto.randomUUID(),
    resources: [],
    verbs: [],
  };
}

/**
 * Converts a role to its in-editor UI model representation. The resulting
 * model may be marked as requiring reset if the role contains unsupported
 * features.
 */
export function roleToRoleEditorModel(
  role: Role,
  originalRole?: Role
): RoleEditorModel {
  // We use destructuring to strip fields from objects and assert that nothing
  // has been left. Therefore, we don't want Lint to warn us that we didn't use
  // some of the fields.
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const { kind, metadata, spec, version, ...unsupported } = role;
  const { name, description, revision, ...unsupportedMetadata } = metadata;
  const { allow, deny, options, ...unsupportedSpecs } = spec;
  const {
    accessSpecs,
    rules,
    requiresReset: allowRequiresReset,
  } = roleConditionsToModel(allow);
  const { model: optionsModel, requiresReset: optionsRequireReset } =
    optionsToModel(options);

  return {
    metadata: {
      name,
      description,
      revision: originalRole?.metadata?.revision,
    },
    accessSpecs,
    rules,
    options: optionsModel,
    requiresReset:
      revision !== originalRole?.metadata?.revision ||
      version !== roleVersion ||
      !(
        isEmpty(unsupported) &&
        isEmpty(unsupportedMetadata) &&
        isEmpty(unsupportedSpecs) &&
        isEmpty(deny)
      ) ||
      allowRequiresReset ||
      optionsRequireReset,
  };
}

/**
 * Converts a `RoleConditions` instance (an "allow" or "deny" section, to be
 * specific) to a list of access specification models.
 */
function roleConditionsToModel(
  conditions: RoleConditions
): Pick<RoleEditorModel, 'accessSpecs' | 'rules' | 'requiresReset'> {
  const {
    node_labels,
    logins,

    kubernetes_groups,
    kubernetes_labels,
    kubernetes_resources,

    app_labels,
    aws_role_arns,
    azure_identities,
    gcp_service_accounts,

    db_labels,
    db_names,
    db_users,
    db_roles,

    windows_desktop_labels,
    windows_desktop_logins,

    rules,

    ...unsupportedConditions
  } = conditions;

  const accessSpecs: AccessSpec[] = [];

  const nodeLabelsModel = labelsToModel(node_labels);
  const nodeLoginsModel = stringsToOptions(logins ?? []);
  if (someNonEmpty(nodeLabelsModel, nodeLoginsModel)) {
    accessSpecs.push({
      kind: 'node',
      labels: nodeLabelsModel,
      logins: nodeLoginsModel,
    });
  }

  const kubeGroupsModel = stringsToOptions(kubernetes_groups ?? []);
  const kubeLabelsModel = labelsToModel(kubernetes_labels);
  const {
    model: kubeResourcesModel,
    requiresReset: kubernetesResourcesRequireReset,
  } = kubernetesResourcesToModel(kubernetes_resources);
  if (someNonEmpty(kubeGroupsModel, kubeLabelsModel, kubeResourcesModel)) {
    accessSpecs.push({
      kind: 'kube_cluster',
      groups: kubeGroupsModel,
      labels: kubeLabelsModel,
      resources: kubeResourcesModel,
    });
  }

  const appLabelsModel = labelsToModel(app_labels);
  const awsRoleARNsModel = aws_role_arns ?? [];
  const azureIdentitiesModel = azure_identities ?? [];
  const gcpServiceAccountsModel = gcp_service_accounts ?? [];
  if (
    someNonEmpty(
      appLabelsModel,
      awsRoleARNsModel,
      azureIdentitiesModel,
      gcpServiceAccountsModel
    )
  ) {
    accessSpecs.push({
      kind: 'app',
      labels: appLabelsModel,
      awsRoleARNs: awsRoleARNsModel,
      azureIdentities: azureIdentitiesModel,
      gcpServiceAccounts: gcpServiceAccountsModel,
    });
  }

  const dbLabelsModel = labelsToModel(db_labels);
  const dbNamesModel = db_names ?? [];
  const dbUsersModel = db_users ?? [];
  const dbRolesModel = db_roles ?? [];
  if (someNonEmpty(dbLabelsModel, dbNamesModel, dbUsersModel, dbRolesModel)) {
    accessSpecs.push({
      kind: 'db',
      labels: dbLabelsModel,
      names: stringsToOptions(dbNamesModel),
      users: stringsToOptions(dbUsersModel),
      roles: stringsToOptions(dbRolesModel),
    });
  }

  const windowsDesktopLabelsModel = labelsToModel(windows_desktop_labels);
  const windowsDesktopLoginsModel = stringsToOptions(
    windows_desktop_logins ?? []
  );
  if (someNonEmpty(windowsDesktopLabelsModel, windowsDesktopLoginsModel)) {
    accessSpecs.push({
      kind: 'windows_desktop',
      labels: windowsDesktopLabelsModel,
      logins: windowsDesktopLoginsModel,
    });
  }

  const { model: rulesModel, requiresReset: rulesRequireReset } =
    rulesToModel(rules);

  return {
    accessSpecs,
    rules: rulesModel,
    requiresReset:
      kubernetesResourcesRequireReset ||
      rulesRequireReset ||
      !isEmpty(unsupportedConditions),
  };
}

function someNonEmpty(...arr: any[][]): boolean {
  return arr.some(x => x.length > 0);
}

/**
 * Converts a set of labels, as represented in the role resource, to a list of
 * `LabelInput` value models.
 */
export function labelsToModel(labels: Labels | undefined): UILabel[] {
  if (!labels) return [];
  return Object.entries(labels).flatMap(([name, value]) => {
    if (typeof value === 'string') {
      return {
        name,
        value,
      };
    } else {
      return value.map(v => ({ name, value: v }));
    }
  });
}

function stringToOption<T extends string>(s: T): Option<T> {
  return { label: s, value: s };
}

function stringsToOptions<T extends string>(arr: T[]): Option<T>[] {
  return arr.map(stringToOption);
}

function kubernetesResourcesToModel(
  resources: KubernetesResource[] | undefined
): { model: KubernetesResourceModel[]; requiresReset: boolean } {
  const result = (resources ?? []).map(kubernetesResourceToModel);
  return {
    model: result.map(r => r.model).filter(m => m !== undefined),
    requiresReset: result.some(r => r.requiresReset),
  };
}

function kubernetesResourceToModel(res: KubernetesResource): {
  model?: KubernetesResourceModel;
  requiresReset: boolean;
} {
  const { kind, name, namespace = '', verbs = [], ...unsupported } = res;
  const kindOption = kubernetesResourceKindOptionsMap.get(kind);
  const verbOptions = verbs.map(verb => kubernetesVerbOptionsMap.get(verb));
  const knownVerbOptions = verbOptions.filter(v => v !== undefined);
  return {
    model:
      kindOption !== undefined
        ? {
            id: crypto.randomUUID(),
            kind: kindOption,
            name,
            namespace,
            verbs: knownVerbOptions,
          }
        : undefined,
    requiresReset:
      kindOption === undefined ||
      verbOptions.length !== knownVerbOptions.length ||
      !isEmpty(unsupported),
  };
}

function rulesToModel(rules: Rule[]): {
  model: RuleModel[];
  requiresReset: boolean;
} {
  const result = (rules ?? []).map(ruleToModel);
  return {
    model: result.map(r => r.model),
    requiresReset: result.some(r => r.requiresReset),
  };
}

function ruleToModel(rule: Rule): { model: RuleModel; requiresReset: boolean } {
  const { resources = [], verbs = [], ...unsupported } = rule;
  const resourcesModel = resources.map(k => resourceKindOptionsMap.get(k));
  const knownResourcesModel = resourcesModel.filter(m => m !== undefined);
  const verbsModel = verbs.map(v => verbOptionsMap.get(v));
  const knownVerbsModel = verbsModel.filter(m => m !== undefined);
  const requiresReset =
    !isEmpty(unsupported) ||
    knownResourcesModel.length !== resourcesModel.length ||
    knownVerbsModel.length !== verbs.length;
  return {
    model: {
      id: crypto.randomUUID(),
      resources: knownResourcesModel,
      verbs: knownVerbsModel,
    },
    requiresReset,
  };
}

function optionsToModel(options: RoleOptions): {
  model: OptionsModel;
  requiresReset: boolean;
} {
  const {
    // Customizable options.
    max_session_ttl,
    client_idle_timeout = '',
    disconnect_expired_cert = false,
    require_session_mfa = false,
    create_host_user_mode = '',
    create_db_user,
    create_db_user_mode = '',
    desktop_clipboard,
    create_desktop_user,
    desktop_directory_sharing,

    // These options must keep their default values, as we don't support them
    // in the standard editor.
    cert_format,
    enhanced_recording,
    forward_agent,
    idp,
    pin_source_ip,
    port_forwarding,
    record_session,
    ssh_file_copy,

    ...unsupported
  } = options;

  const requireMFATypeOption =
    requireMFATypeOptionsMap.get(require_session_mfa);
  const createHostUserModeOption = createHostUserModeOptionsMap.get(
    create_host_user_mode
  );
  const createDBUserModeOption =
    createDBUserModeOptionsMap.get(create_db_user_mode);

  const defaultOpts = defaultOptions();

  return {
    model: {
      maxSessionTTL: max_session_ttl,
      clientIdleTimeout: client_idle_timeout,
      disconnectExpiredCert: disconnect_expired_cert,
      requireMFAType:
        requireMFATypeOption ?? requireMFATypeOptionsMap.get(false),
      createHostUserMode:
        createHostUserModeOption ?? createHostUserModeOptionsMap.get(''),
      createDBUser: create_db_user,
      createDBUserMode:
        createDBUserModeOption ?? createDBUserModeOptionsMap.get(''),
      desktopClipboard: desktop_clipboard,
      createDesktopUser: create_desktop_user,
      desktopDirectorySharing: desktop_directory_sharing,
    },

    requiresReset:
      cert_format !== defaultOpts.cert_format ||
      !equalsDeep(enhanced_recording, defaultOpts.enhanced_recording) ||
      forward_agent !== defaultOpts.forward_agent ||
      !equalsDeep(idp, defaultOpts.idp) ||
      pin_source_ip !== defaultOpts.pin_source_ip ||
      port_forwarding !== defaultOpts.port_forwarding ||
      !equalsDeep(record_session, defaultOpts.record_session) ||
      ssh_file_copy !== defaultOpts.ssh_file_copy ||
      requireMFATypeOption === undefined ||
      createHostUserModeOption === undefined ||
      createDBUserModeOption === undefined ||
      !isEmpty(unsupported),
  };
}

function isEmpty(obj: object) {
  return Object.keys(obj).length === 0;
}

/**
 * Converts a role editor model to a role. This operation is lossless.
 */
export function roleEditorModelToRole(roleModel: RoleEditorModel): Role {
  const { name, description, revision, ...mRest } = roleModel.metadata;
  // Compile-time assert that protects us from silently losing fields.
  mRest satisfies Record<any, never>;

  const role: Role = {
    kind: 'role',
    metadata: {
      name,
      description,
      revision,
    },
    spec: {
      allow: {},
      deny: {},
      options: optionsModelToRoleOptions(roleModel.options),
    },
    version: roleVersion,
  };

  for (const spec of roleModel.accessSpecs) {
    const { kind } = spec;
    switch (kind) {
      case 'node':
        role.spec.allow.node_labels = labelsModelToLabels(spec.labels);
        role.spec.allow.logins = optionsToStrings(spec.logins);
        break;

      case 'kube_cluster':
        role.spec.allow.kubernetes_groups = optionsToStrings(spec.groups);
        role.spec.allow.kubernetes_labels = labelsModelToLabels(spec.labels);
        role.spec.allow.kubernetes_resources = spec.resources.map(
          ({ kind, name, namespace, verbs }) => ({
            kind: kind.value,
            name,
            namespace,
            verbs: optionsToStrings(verbs),
          })
        );
        break;

      case 'app':
        role.spec.allow.app_labels = labelsModelToLabels(spec.labels);
        role.spec.allow.aws_role_arns = spec.awsRoleARNs;
        role.spec.allow.azure_identities = spec.azureIdentities;
        role.spec.allow.gcp_service_accounts = spec.gcpServiceAccounts;
        break;

      case 'db':
        role.spec.allow.db_labels = labelsModelToLabels(spec.labels);
        role.spec.allow.db_names = optionsToStrings(spec.names);
        role.spec.allow.db_users = optionsToStrings(spec.users);
        role.spec.allow.db_roles = optionsToStrings(spec.roles);
        break;

      case 'windows_desktop':
        role.spec.allow.windows_desktop_labels = labelsModelToLabels(
          spec.labels
        );
        role.spec.allow.windows_desktop_logins = optionsToStrings(spec.logins);
        break;

      default:
        kind satisfies never;
    }
  }

  if (roleModel.rules.length > 0) {
    role.spec.allow.rules = roleModel.rules.map(role => ({
      resources: role.resources.map(r => r.value),
      verbs: role.verbs.map(v => v.value),
    }));
  }

  return role;
}

/**
 * Converts a list of `LabelInput` value models to a set of labels, as
 * represented in the role resource.
 */
export function labelsModelToLabels(uiLabels: UILabel[]): Labels {
  const labels = {};
  for (const { name, value } of uiLabels) {
    if (!Object.hasOwn(labels, name)) {
      labels[name] = value;
    } else if (typeof labels[name] === 'string') {
      labels[name] = [labels[name], value];
    } else {
      labels[name].push(value);
    }
  }
  return labels;
}

function optionsModelToRoleOptions(model: OptionsModel): RoleOptions {
  return {
    ...defaultOptions(),

    // Note: technically, coercing the optional fields to undefined is not
    // necessary, but it's easier to test it this way, since we achieve
    // symmetry between what goes into the model and what goes out of it, even
    // if some fields are optional.
    max_session_ttl: model.maxSessionTTL,
    client_idle_timeout: model.clientIdleTimeout || undefined,
    disconnect_expired_cert: model.disconnectExpiredCert || undefined,
    require_session_mfa: model.requireMFAType.value || undefined,
    create_host_user_mode: model.createHostUserMode.value || undefined,
    create_db_user: model.createDBUser,
    create_db_user_mode: model.createDBUserMode.value || undefined,
    desktop_clipboard: model.desktopClipboard,
    create_desktop_user: model.createDesktopUser,
    desktop_directory_sharing: model.desktopDirectorySharing,
  };
}

function optionsToStrings<T = string>(opts: readonly Option<T>[]): T[] {
  return opts.map(opt => opt.value);
}

/** Detects if fields were modified by comparing against the original role. */
export function hasModifiedFields(
  updated: RoleEditorModel,
  originalRole: Role
) {
  return !equalsDeep(roleEditorModelToRole(updated), originalRole, {
    ignoreUndefined: true,
  });
}
