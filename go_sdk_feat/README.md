

ls /Users/bytedance/work/go/pkg/mod/github.com/olivere/elastic/v7@v7.0.22

ls $(GOPATH) + /pkg/mod/ + github.com/olivere/elastic/v7@v7.0.22



ls /Users/bytedance/work/go/pkg/mod/github.com/olivere/elastic/v7@v7.0.22


find . /Users/bytedance/work/go/pkg/mod/github.com/olivere/elastic/v7@v7.0.22 | grep "\.go" | grep -v "_test"

find . /Users/bytedance/work/go/pkg/mod/github.com/olivere/elastic/v7@v7.0.22 | grep "\.go" | grep -v "_test" | xargs cat | grep "func New" > go_sdk_feat/all_feat.txt

cat go_sdk_feat/all_feat.txt | grep "Query(" > go_sdk_feat/all_feat_query.txt
cat all_feat.txt | grep "Aggregation(" > all_feat_aggregation.txt
cat all_feat.txt | grep "Function(" > all_feat_function.txt

cat all_feat.txt | grep "Service(" > all_feat_service.txt
cat all_feat.txt | grep "Client(" > all_feat_client.txt

cat all_feat.txt | grep "Model(" > all_feat_model.txt
cat all_feat.txt | grep "Sort(" > all_feat_sort.txt




find . | grep txt | xargs cat | sort | uniq -c | grep  "1 func " | grep "Service"
find . | grep txt | xargs cat | sort | uniq -c | grep  "1 func " | grep "Client("
