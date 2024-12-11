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

import React from 'react';
import Table from 'design/DataTable';

import { useParams } from 'react-router';

import { FeatureBox } from 'teleport/components/Layout';
import { AwsOidcHeader } from 'teleport/Integrations/status/AwsOidc/AwsOidcHeader';
import { AwsResource } from 'teleport/Integrations/status/AwsOidc/StatCard';
import { useAwsOidcStatus } from 'teleport/Integrations/status/AwsOidc/useAwsOidcStatus';
import { IntegrationKind } from 'teleport/services/integrations';

export function Tasks() {
  const { integrationAttempt } = useAwsOidcStatus();
  const { data: integration } = integrationAttempt;

  return (
    <FeatureBox css={{ maxWidth: '1400px', paddingTop: '16px', gap: '30px' }}>
      {integration && (
        <AwsOidcHeader integration={integration} tasks={true} />
      )}
      <Table
        data={[]}
        columns={[
          {
            key: 'type',
            headerText: 'Type',
            isSortable: true,
          },
          {
            key: 'details',
            headerText: 'Issue Details',
            isSortable: true,
          },
          {
            key: 'timestamp',
            headerText: 'Timestamp',
            isSortable: true,
          },
        ]}
        emptyText={`No pending tasks`}
        isSearchable
      />
    </FeatureBox>
  );
}
