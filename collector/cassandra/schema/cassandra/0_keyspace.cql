CREATE KEYSPACE IF NOT EXISTS otel
  WITH REPLICATION = {
   'class' : 'SimpleStrategy',
   'replication_factor' : 1
  };