var addon = require("bindings")("hello");

module.exports.message = addon.hello;
