import { Component, type OpenApiV3_1 } from "fluid-oas";
import { ERROR_SCHEMA } from "./error";

const COMPONENT = Component.addSchemas({
  ErrorSchema: ERROR_SCHEMA,
});

export const COMPONENT_NAME_MAPPINGS = COMPONENT.createMappings();

export function addOpenApiComponents(oas: OpenApiV3_1): OpenApiV3_1 {
  return oas.addComponents(COMPONENT);
}
