---
title: mdir2mailbox.pl
---

```perl
#!/usr/bin/env perl
# dw-maildirtombox.pl
# dw = Dovecot Wiki :-)
# NOTE! The output file must not contain single quotes (')!
# figure out which program to run
$cmd="reformail -f1";
system("$cmd </dev/null >/dev/null 2>/dev/null") == 0 or $cmd="formail";
system("$cmd </dev/null >/dev/null 2>/dev/null") == 0
or die "cannot find reformail or formail on your \$PATH!\nAborting";
$dir=$ARGV[0];
$outputfile=$ARGV[1];
if (($outputfile eq '') || ($dir eq ''))
{ die "Usage: ./archivemail.pl mailbox outputfile\nAborting"; }
if (!stat("Maildir/$dir/cur") || !stat("Maildir/$dir/new"))
{ die "Maildir/$dir is not a maildir.\nAborting"; }
@files = (<Maildir/$dir/cur/*>,<Maildir/$dir/new/*>);
foreach $file (@files) {
  next unless -f $file; # skip non-regular files
  next unless -s $file; # skip empty files
  next unless -r $file; # skip unreadable files
  $file =~ s/'/'"'"'/;  # escape ' (single quote)
  $run = "cat '$file' | $cmd >>'$outputfile'";
  system($run) == 0 or warn "cannot run \"$run\".";
}
```