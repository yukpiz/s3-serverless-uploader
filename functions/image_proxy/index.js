'use strict';

exports.handler = (event, context, callback) => {
  const request = event.Records[0].cf.request;
  // const headers = request.headers;

  console.log(request);
  callback(null, request);
};
