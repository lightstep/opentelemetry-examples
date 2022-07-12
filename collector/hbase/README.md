[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/big-data-europe/Lobby)

# docker-hbase

# Standalone
To run standalone hbase:
```
docker-compose -f docker-compose.yml up -d
```
The deployment is the same as in [quickstart HBase documentation](https://hbase.apache.org/book.html#quickstart).
Can be used for testing/development, connected to Hadoop cluster.

This deployment will start Zookeeper, HMaster and HRegionserver in separate containers.
