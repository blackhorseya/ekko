import { test, expect } from '@playwright/test';

let id = 0;

test.describe('Tasks Testing', () => {
  test.beforeAll(async ({ request }) => {
    const resp = await request.post('/api/v1/tasks', {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      form: {
        'title': 'title',
      },
    });
    expect(resp.ok()).toBeTruthy();

    const task = await resp.json()
    id = task.data.id;
  });
  
  test.afterAll(async ({ request }) => {
    const resp = await request.delete(`/api/v1/tasks/${id}`);
    expect(resp.ok()).toBeTruthy();
  });

  test('Should get a task by id', async ({request}) => {
    const resp = await request.get(`/api/v1/tasks/${id}`);
    expect(resp.ok()).toBeTruthy();
  });

  test('Should update a task title', async ({ request }) => {
    const resp = await request.patch(`/api/v1/tasks/${id}/title`, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      form: {
        'title': 'title',
      },
    })
    expect(resp.ok()).toBeTruthy();
  });

  test('Should update a task status', async ({ request }) => {
    const resp = await request.patch(`/api/v1/tasks/${id}/status`, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      form: {
        'status': 2,
      },
    })
    expect(resp.ok()).toBeTruthy();
  });
});
