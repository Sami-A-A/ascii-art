echo "Test 1:"
go run . "" | cat -e
echo "Test 2:"
go run . "\n" | cat -e
echo "Test 3:"
go run . "\n\n" | cat -e
echo "Test 4:"
 go run . "\nhello" | cat -e
echo "Test 5:"
 go run . "\nhello\n" | cat -e
echo "Test 6:"
 go run . "hello\n" | cat -e
echo "Test 7:"
 go run . "Hello\nThere" | cat -e
echo "Test 8:"
 go run . "Hello\n\nThere" | cat -e
