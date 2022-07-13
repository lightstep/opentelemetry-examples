// -------------- otelu_metrics ----
DROP TABLE IF EXISTS otel.otelu_metrics;
DROP TYPE IF EXISTS otel.metric;

CREATE TYPE otel.metric (
   metric_title text,
   metric_date timestamp,
   metric_time time);

CREATE TABLE otel.otelu_metrics (
  id UUID PRIMARY KEY,
  lastname text,
  firstname text,
  metrics list<FROZEN <metric>> );

INSERT INTO otel.otelu_metrics (id,metrics) VALUES (
   e7ae5cf3-d358-4d99-b900-85902fda9bb0,
   [ { metric_title:'17th Santos Tour Down Under Aalburg', metric_date:'2017-04-14',metric_time:'07:00:00' },
     { metric_title:'17th Santos Tour Down Under Gelderland', metric_date:'2017-04-14', metric_time:'08:00:00' } ]);

INSERT INTO otel.otelu_metrics (id, lastname, firstname, metrics) 
VALUES (
	e7cd5752-bc0d-4157-a80f-7523add8dbcd, 'VAN DER BREGGEN', 'Anna', 
	[ {metric_title:'Festival Luxembourgeois du otelme feminin Elsy Jacobs - Prologue - Garnich > Garnich',metric_date:'2017-04-14',metric_time:'08:00:00'},
		{metric_title:'Festival Luxembourgeois du otelme feminin Elsy Jacobs - Stage 2 - Garnich > Garnich',metric_date:'2017-04-14',metric_time:'06:00:00'},
		{metric_title:'Festival Luxembourgeois du otelme feminin Elsy Jacobs - Stage 3 - Mamer > Mamer',metric_date:'2017-04-14',metric_time:'10:00:00'} ]);

UPDATE otel.otelu_metrics
	SET 
    lastname = 'FRAME',
    firstname = 'Alex',
    metrics[1] = { metric_time:'06:00:00'}
		WHERE id = e7ae5cf3-d358-4d99-b900-85902fda9bb0;