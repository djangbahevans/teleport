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

import React, { useState } from 'react';
import { ButtonPrimary, Text, Box, ButtonIcon, Menu } from 'design';
import { Info } from 'design/Icon';
import { format } from 'date-fns';
import { ResourceIdKind } from 'teleport/services/agents';

import { HoverTooltip } from 'shared/components/ToolTip';
import cfg from 'shared/config';

import { AccessRequest } from 'shared/services/accessRequests';

export function PromotedMessage({
  request,
  px,
  py,
  self,
  assumeAccessList,
}: {
  request: AccessRequest;
  self: boolean;
  px?: number;
  py?: number;
  assumeAccessList(): void;
}) {
  const { promotedAccessListTitle, user } = request;

  return (
    <Box px={px} py={py}>
      <Text>
        This access request has been promoted to long-term access.
        <br />
        {self ? (
          <>
            You are now a member of Access List <b>{promotedAccessListTitle}</b>{' '}
            which grants you the resources requested.
          </>
        ) : (
          <>
            {user} is now a member of Access List{' '}
            <b>{promotedAccessListTitle}</b> which grants {user} the resources
            requested.
          </>
        )}
      </Text>
      {self && (
        <ButtonPrimary mt={3} onClick={assumeAccessList}>
          Re-login to gain access
        </ButtonPrimary>
      )}
    </Box>
  );
}

export const ButtonPromotedInfo = ({
  request,
  ownRequest,
  assumeAccessList,
}: {
  request: AccessRequest;
  ownRequest: boolean;
  assumeAccessList(): void;
}) => {
  const [anchorEl, setAnchorEl] = useState(null);

  const handleOpen = event => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <Box css={{ margin: '0 auto' }}>
      <ButtonIcon onClick={handleOpen}>
        <Info />
      </ButtonIcon>
      <Menu
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        transformOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        <PromotedMessage
          request={request}
          self={ownRequest}
          assumeAccessList={assumeAccessList}
          px={4}
          py={4}
        />
      </Menu>
    </Box>
  );
};

export function getAssumeStartTimeTooltipText(startTime: Date) {
  const formattedDate = format(startTime, cfg.dateWithPrefixedTime);
  return `Access is not available until the approved time of ${formattedDate}`;
}

export const BlockedByStartTimeButton = ({
  assumeStartTime,
}: {
  assumeStartTime: Date;
}) => {
  return (
    <HoverTooltip
      tipContent={getAssumeStartTimeTooltipText(assumeStartTime)}
      anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
      transformOrigin={{ vertical: 'bottom', horizontal: 'right' }}
    >
      <ButtonPrimary disabled={true} size="small">
        Assume Roles
      </ButtonPrimary>
    </HoverTooltip>
  );
};

/** Available request kinds for resource-based and role-based access requests. */
export type ResourceKind = ResourceIdKind | 'role' | 'resource';

export function getPrettyResourceKind(kind: ResourceKind): string {
  switch (kind) {
    case 'role':
      return 'Role';
    case 'app':
      return 'Application';
    case 'node':
      return 'Server';
    case 'resource':
      return 'Resource';
    case 'db':
      return 'Database';
    case 'kube_cluster':
      return 'Kubernetes';
    case 'user_group':
      return 'User Group';
    case 'windows_desktop':
      return 'Desktop';
    case 'saml_idp_service_provider':
      return 'SAML Application';
    case 'aws_iam_ic_account':
      return 'AWS Identity Center Account';
    case 'aws_iam_ic_account_assignment':
      return 'AWS Account Assignment';
    default:
      kind satisfies never;
      return kind;
  }
}
