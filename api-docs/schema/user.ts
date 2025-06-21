import { Object, String } from "fluid-oas";
import { ID } from "../paths/utils";

const userRequired = ["firstName", "lastName", "username", "email"];

export const USER_REQUEST_SCHEMA = Object.addProperties({
  firstName: String.addMinLength(1),
  lastName: String.addMinLength(1),
  username: String.addMinLength(3),
  email: String.addFormat("email"),
})
  .addDescription("User schema represents a user in Hangouts.")
  .addExample({
    firstName: "John",
    lastName: "Smith",
    username: "JohnSlayer69",
    email: "JohnSlayer69@blah.com",
  })
  .addRequired(userRequired);

export const USER_SCHEMA = Object.addProperties({
  firstName: String,
  lastName: String,
  username: String,
  email: String,
  id: ID,
}).addRequired(userRequired.concat(["id"]));
