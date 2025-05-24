module.exports = grammar({
  name: "authzed",

  rules: {
    source_file: ($) => repeat($.definition),

    definition: ($) =>
      seq("definition", $.identifier, "{", repeat($.statement), "}"),

    statement: ($) => choice($.relation, $.permission),

    relation: ($) => seq("relation", $.identifier, ":", $.type),

    permission: ($) => seq("permission", $.identifier, "=", $.expression),

    type: ($) => $.identifier,

    expression: ($) =>
      choice($.identifier, seq($.identifier, ".", $.identifier)),

    identifier: ($) => /[a-zA-Z_][a-zA-Z0-9_]*/,
  },
});
