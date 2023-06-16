/**
 * Bahn Alarm
 * An app for live tracking of train delays from the DB  - In case of a request validation failure a `400` status code is returned regardless of what\'s documented - All endpoints that document a `401` response require authentication   - If the JWT is expired or invalid a `401` status code is returned   - If the `jwt` cookie is not present a `400` status code is returned
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

export interface VapidKeys {
  publicKey: string;
}
