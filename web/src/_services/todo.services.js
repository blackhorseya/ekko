function list(start, end) {
  const opts = {
    method: 'GET',
  };

  return fetch(
      `${process.env.REACT_APP_API_ENDPOINT}/api/v1/tasks？start=${start}&end=${end}`,
      opts).
      then(handlerResp);
}

function handlerResp(response) {
  return response.text().then(text => {
    console.log(text);
    const resp = text && JSON.parse(text);

    if (!response.ok) {
      const error = (resp && resp.msg) || response.statusText;
      return Promise.reject(error);
    }

    return resp.data;
  });
}