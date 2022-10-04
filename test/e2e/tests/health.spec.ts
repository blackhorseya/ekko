import { test, expect } from '@playwright/test';

test.describe('Health Testing', () => {
  test('Readiness should be ok', async ({ request }) => {
    const resp = await request.get(`/api/readiness`);

    expect(resp.ok()).toBeTruthy();
  });

  test('Liveness should be ok', async ({ request }) => {
    const resp = await request.get(`/api/liveness`);

    expect(resp.ok()).toBeTruthy();
  });
});
