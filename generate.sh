for i in `find . -name action.yaml | sed 's|/[^/]*$||'`; do
    cat <<EOF >$i/README.md
<!-- action-docs-description -->

<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

<!-- action-docs-runs -->
EOF
done
