// source: teleport/lib/teleterm/vnet/v1/vnet_service.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

goog.exportSymbol('proto.teleport.lib.teleterm.vnet.v1.StartRequest', null, global);
goog.exportSymbol('proto.teleport.lib.teleterm.vnet.v1.StartResponse', null, global);
goog.exportSymbol('proto.teleport.lib.teleterm.vnet.v1.StopRequest', null, global);
goog.exportSymbol('proto.teleport.lib.teleterm.vnet.v1.StopResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.teleport.lib.teleterm.vnet.v1.StartRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.teleport.lib.teleterm.vnet.v1.StartRequest.displayName = 'proto.teleport.lib.teleterm.vnet.v1.StartRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.teleport.lib.teleterm.vnet.v1.StartResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.teleport.lib.teleterm.vnet.v1.StartResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.teleport.lib.teleterm.vnet.v1.StartResponse.displayName = 'proto.teleport.lib.teleterm.vnet.v1.StartResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.teleport.lib.teleterm.vnet.v1.StopRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.teleport.lib.teleterm.vnet.v1.StopRequest.displayName = 'proto.teleport.lib.teleterm.vnet.v1.StopRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.teleport.lib.teleterm.vnet.v1.StopResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.teleport.lib.teleterm.vnet.v1.StopResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.teleport.lib.teleterm.vnet.v1.StopResponse.displayName = 'proto.teleport.lib.teleterm.vnet.v1.StopResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.teleport.lib.teleterm.vnet.v1.StartRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StartRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    rootClusterUri: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StartRequest}
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.teleport.lib.teleterm.vnet.v1.StartRequest;
  return proto.teleport.lib.teleterm.vnet.v1.StartRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StartRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StartRequest}
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRootClusterUri(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.teleport.lib.teleterm.vnet.v1.StartRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StartRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRootClusterUri();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string root_cluster_uri = 1;
 * @return {string}
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.prototype.getRootClusterUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StartRequest} returns this
 */
proto.teleport.lib.teleterm.vnet.v1.StartRequest.prototype.setRootClusterUri = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.teleport.lib.teleterm.vnet.v1.StartResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.teleport.lib.teleterm.vnet.v1.StartResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StartResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StartResponse.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StartResponse}
 */
proto.teleport.lib.teleterm.vnet.v1.StartResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.teleport.lib.teleterm.vnet.v1.StartResponse;
  return proto.teleport.lib.teleterm.vnet.v1.StartResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StartResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StartResponse}
 */
proto.teleport.lib.teleterm.vnet.v1.StartResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.teleport.lib.teleterm.vnet.v1.StartResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.teleport.lib.teleterm.vnet.v1.StartResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StartResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StartResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.teleport.lib.teleterm.vnet.v1.StopRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StopRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    rootClusterUri: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StopRequest}
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.teleport.lib.teleterm.vnet.v1.StopRequest;
  return proto.teleport.lib.teleterm.vnet.v1.StopRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StopRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StopRequest}
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRootClusterUri(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.teleport.lib.teleterm.vnet.v1.StopRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StopRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRootClusterUri();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string root_cluster_uri = 1;
 * @return {string}
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.prototype.getRootClusterUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StopRequest} returns this
 */
proto.teleport.lib.teleterm.vnet.v1.StopRequest.prototype.setRootClusterUri = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.teleport.lib.teleterm.vnet.v1.StopResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.teleport.lib.teleterm.vnet.v1.StopResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StopResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StopResponse.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StopResponse}
 */
proto.teleport.lib.teleterm.vnet.v1.StopResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.teleport.lib.teleterm.vnet.v1.StopResponse;
  return proto.teleport.lib.teleterm.vnet.v1.StopResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StopResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.teleport.lib.teleterm.vnet.v1.StopResponse}
 */
proto.teleport.lib.teleterm.vnet.v1.StopResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.teleport.lib.teleterm.vnet.v1.StopResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.teleport.lib.teleterm.vnet.v1.StopResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.teleport.lib.teleterm.vnet.v1.StopResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.teleport.lib.teleterm.vnet.v1.StopResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


goog.object.extend(exports, proto.teleport.lib.teleterm.vnet.v1);
