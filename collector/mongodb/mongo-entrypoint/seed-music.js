print('===============JAVASCRIPT===============');
print('Count of rows in test collection: ' + db.test.count());

db.test.insert({ myfield: 'test1', anotherfield: 'TEST1' });
db.test.insert({ myfield: 'test2', anotherfield: 'TEST2' });

print('===============AFTER JS INSERT==========');
print('Count of rows in test collection: ' + db.test.count());

alltest = db.test.find();
while (alltest.hasNext()) {
  printjson(alltest.next());
}