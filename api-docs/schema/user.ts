import { Integer, Object, String } from "fluid-oas";
import { ID } from "../paths/utils";

export const USER_REQUEST_SCHEMA = Object.addProperties({
  firstName: String.addMinLength(1),
  lastName: String.addMinLength(1),
  username: String.addMinLength(3),
  age: Integer.addMinimum(0),
})
  .addDescription("User schema represents a user in Hangouts.")
  .addExample({
    firstName: "John",
    lastName: "Smith",
    username: "JohnSlayer69",
    age: 18,
  })
  .addRequired(["firstName", "lastName", "username", "age"]);

export const USER_SCHEMA = Object.addProperties({
  firstName: String,
  lastName: String,
  username: String,
  age: Integer,
  id: ID,
}).addRequired(["firstName", "lastName", "username", "age", "id"]);
