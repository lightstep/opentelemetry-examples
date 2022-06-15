print('===============JAVASCRIPT===============');
print('Count of rows in otel collection: ' + db.otel.count());

db.otel.insert({ fullName: 'test1', age: '30' });
db.otel.insert({ fullName: 'test2', age: '23' });

print('===============AFTER JS INSERT==========');
print('Count of rows in otel collection: ' + db.otel.count());

allotel = db.otel.find();
while (allotel.hasNext()) {
  printjson(allotel.next());
}
