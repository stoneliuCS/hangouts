import { Object, String } from "fluid-oas";

export const ERROR_SCHEMA = Object.addProperties({
  error: String,
})
  .addDescription("Error Schema.")
  .addRequired(["error"]);
