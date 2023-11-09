# paramremove

go install github.com/erickfernandox/paramremove@latest

How to work


Urls with parameters:

URL: htttps://example.com?x=x&y=y

echo "htttps://example.com?x=x&y=y"|paramremove
cat list_urls.txt | paramremove

result:

htttps://example.com
