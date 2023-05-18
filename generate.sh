pwd=`pwd`
for i in `find . -name action.yaml | sed 's|/[^/]*$||'`; do
    cat <<EOF >$pwd/$i/README.md
<!-- action-docs-description -->

<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
EOF
cd $pwd/$i
action-docs --action ./action.yaml --update-readme
done
