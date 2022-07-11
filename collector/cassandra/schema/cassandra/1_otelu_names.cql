// -------------- otelu_name ---
DROP TABLE IF EXISTS otel.otelu_name;

CREATE TABLE otel.otelu_name (
  id UUID PRIMARY KEY,
  lastname text,
  firstname text
);

// Insert a record that only contains the min values UUID
INSERT INTO otel.otelu_name (id) VALUES (uuid());

// Remove the record from the table
TRUNCATE otel.otelu_name;

INSERT INTO otel.otelu_name (id, lastname, firstname) VALUES (e7cd5752-bc0d-4157-a80f-7523add8dbcd, 'VAN DER BREGGEN', 'Anna');
INSERT INTO otel.otelu_name (id, lastname, firstname) VALUES (e7ae5cf3-d358-4d99-b900-85902fda9bb0, 'FRAME', 'Alex');
INSERT INTO otel.otelu_name (id, lastname, firstname) VALUES (220844bf-4860-49d6-9a4b-6b5d3a79cbfb, 'TIRALONGO', 'Paolo');
INSERT INTO otel.otelu_name (id, lastname, firstname) VALUES (6ab09bec-e68e-48d9-a5f8-97e6fb4c9b47, 'KRUIKSWIJK', 'Steven');
INSERT INTO otel.otelu_name (id, lastname, firstname) VALUES (fb372533-eb95-4bb4-8685-6ef61e994caa, 'MATTHEWS', 'Michael');