# Set version
tag=$1
: > ./cmd/VERSION && echo $tag >  ./cmd/VERSION

# Commit version number & push
git add ./cmd/VERSION
git commit -m "Bump version to $tag"
git push origin

# Tag & Push.
git tag $tag
git push origin $tag
