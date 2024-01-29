echo "Test 1:"
go run . "" | cat -e
echo "Test 2:"
go run . "\n" | cat -e
echo "Test 3:"
go run . "\n\n" | cat -e
echo "Test 4:"
 go run . "\nhello" | cat -e
echo "Test 5:"
<<<<<<< HEAD:main/test.sh
 go run . "\nhello\n" | cat -e
echo "Test 6:"
 go run . "hello\n" | cat -e
=======
 go run . "\nHeLlO" | cat -e
echo "Test 6:"
 go run . "Hello\n\n" | cat -e
>>>>>>> 6e95d2dde7a60bb4680c33a9a990fde3cd4c4ebf:test/test.sh
echo "Test 7:"
 go run . "Hello\nThere" | cat -e
echo "Test 8:"
 go run . "Hello\n\nThere" | cat -e
