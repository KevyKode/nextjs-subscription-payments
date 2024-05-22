import { powerbiConfig } from './powerbi.config';

const apiUrl = 'https://api.powerbi.com/v1.0';

const authenticate = async () => {
  const authUrl = `https://login.microsoftonline.com/${powerbiConfig.tenantId}/oauth2/v2.0/token`;
  const headers = {
    'Content-Type': 'application/x-www-form-urlencoded',
  };
  const data = `grant_type=client_credentials&client_id=${powerbiConfig.clientId}&client_secret=${powerbiConfig.clientSecret}`;

  const response = await fetch(authUrl, {
    method: 'POST',
    headers,
    body: data,
  });

  const tokenResponse = await response.json();
  return tokenResponse.access_token;
};

const getDashboard = async (accessToken: string) => {
  const dashboardUrl = `${apiUrl}/groups/${powerbiConfig.groupId}/dashboards/${powerbiConfig.dashboardId}`;
  const headers = {
    Authorization: `Bearer ${accessToken}`,
  };

  const response = await fetch(dashboardUrl, {
    method: 'GET',
    headers,
  });

  const dashboardResponse = await response.json();
  return dashboardResponse;
};

export const powerbiApi = {
  authenticate,
  getDashboard,
};