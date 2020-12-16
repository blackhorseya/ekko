export const taskServices = {
  list,
};

function list() {
  const reqOpt = {
    method: 'GET',
  };

  return fetch('http://localhost:5000/api/v1/tasks', reqOpt).
      then(handleResponse);
}

function handleResponse(resp) {
  return resp.text().then(body => {
    const data = body && JSON.parse(body);
    if (resp.status !== 200) {
      const error = (data && data.message) || resp.statusText;
      return Promise.reject(error);
    }

    return data;
  });
}