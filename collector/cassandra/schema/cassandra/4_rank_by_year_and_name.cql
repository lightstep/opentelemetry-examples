// ------------- rank_by_year_and_name 
DROP TABLE IF EXISTS otel.rank_by_year_and_name;

CREATE TABLE otel.rank_by_year_and_name (  metric_year int,
  metric_name text,
  otelu_name text,
  rank int,
  PRIMARY KEY ((metric_year, metric_name), rank)
);

INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2015, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Benjamin PRADES', 1);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2015, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Adam PHELAN', 2);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2015, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Thomas LEBAS', 3);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Benjamin PRADES', 3);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Daniel MARTIN', 1);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, 'Tour of Japan - Stage 4 - Minami > Shinshu', 'Johan Esteban CHAVES', 2);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2015, 'Giro d''Italia - Stage 11 - Forli > Imola', 'Ilnur ZAKARIN', 1);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2015, 'Giro d''Italia - Stage 11 - Forli > Imola', 'Carlos BETANCUR', 2);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, '4th Tour of Beijing', 'Phillippe GILBERT', 1);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, '4th Tour of Beijing', 'Daniel MARTIN', 2);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, '4th Tour of Beijing', 'Johan Esteban CHAVES', 3);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, '4th Tour of Beijing', 'Phillippe GILBERT', 1);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, '4th Tour of Beijing', 'Daniel MARTIN', 2);
INSERT INTO otel.rank_by_year_and_name (  metric_year, metric_name, otelu_name, rank ) VALUES (   2014, '4th Tour of Beijing', 'Johan Esteban CHAVES', 3);