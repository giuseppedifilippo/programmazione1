#split --numeric-suffix -l $1 $2 $2"."
split --numeric-suffix -l $1 $2 taglio-
more taglio*|tee
rm taglio*
