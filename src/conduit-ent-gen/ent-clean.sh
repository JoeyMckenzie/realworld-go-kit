@mkdir ./tmp
@cp -r ./ent/schema ./tmp
@cp ./ent/generate.go ./tmp
@rm -rf ./ent
@mkdir ./ent
@cp -r ./tmp/schema ./ent
@cp ./tmp/generate.go ./ent
@rm -rf ./tmp
