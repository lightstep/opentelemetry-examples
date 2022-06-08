// ------------------- metric_times ---
DROP TABLE IF EXISTS otel.metric_times;

CREATE TABLE otel.metric_times (
    metric_name text,
    metric_time time,
    finish_time duration,
    metric_date date,
    otelu_name text,
    PRIMARY KEY (metric_date, metric_name, otelu_name));

INSERT INTO otel.metric_times (metric_name, otelu_name, metric_time, metric_date) VALUES ('17th Santos Tour Down Under Aalburg', 'Anna VAN DER BREGGEN', '19:15:18', '2017-04-14');
INSERT INTO otel.metric_times (metric_name, otelu_name, metric_time, metric_date) VALUES ('17th Santos Tour Down Under Aalburg', 'Alex FRAME', '19:15:23', '2017-04-14' );
INSERT INTO otel.metric_times (metric_name, otelu_name, metric_time, metric_date) VALUES ('17th Santos Tour Down Under Aalburg', 'Paolo TIRALONGO', '19:15:28', '2017-04-14');
INSERT INTO otel.metric_times (metric_name, otelu_name, metric_time, metric_date) VALUES ('17th Santos Tour Down Under Aalburg', 'Steven KRUIKSWIJK', '19:15:33', '2017-04-14');
INSERT INTO otel.metric_times (metric_name, otelu_name, metric_time, metric_date) VALUES ('17th Santos Tour Down Under Aalburg', 'Michael MATTHEWS', '19:15:38', '2017-04-14' );