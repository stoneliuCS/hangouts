import { Component, SecurityScheme, type OpenApiV3_1 } from "fluid-oas";
import { ERROR_SCHEMA } from "./error";
import { USER_SCHEMA } from "./user";

const COMPONENT = Component.addSchemas({
  ErrorSchema: ERROR_SCHEMA,
  UserSchema: USER_SCHEMA,
}).addSecuritySchemes({
  bearerAuth: SecurityScheme.addType("http")
    .addScheme("bearer")
    .addBearerFormat("JWT"),
});

export function addOpenApiComponents(oas: OpenApiV3_1): OpenApiV3_1 {
  return oas.addComponents(COMPONENT);
}
