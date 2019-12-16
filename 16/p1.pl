#!/usr/bin/env perl

use 5.28.0;
use List::Util qw(sum);

sub generate_factors {
    my ($signal_len) = @_;

    state $base_procs = [ 0, 1, 0, -1 ];

    my @factors;
    for my $phase ( 1 .. $signal_len ) {
        my @fs;

        my $b = 0;
        while (1) {
            if ( scalar @fs > $signal_len ) {
                last;
            }

            push @fs, ( $base_procs->[ $b++ ] ) x $phase;
            if ( $b == scalar $base_procs->@* ) {
                $b = 0;
            }
        }

        push @factors, [ map { $fs[$_] } ( 1 .. $signal_len ) ];
    }

    return \@factors;
}

my $input_signal = <>;
chomp($input_signal);
my @curr_sigs = split //, $input_signal;
my $sig_len   = scalar @curr_sigs;
my $factors   = generate_factors($sig_len);

for my $phase ( 1 .. 100 ) {
    my @next_sigs;

    for my $i ( 0 .. $sig_len - 1 ) {
        push @next_sigs,
          abs( sum map { $curr_sigs[$_] * $factors->[$i]->[$_] }
              ( 0 .. $sig_len - 1 ) ) % 10;
    }

    @curr_sigs = @next_sigs;
}

say join( '', @curr_sigs[ 0 .. 7 ] );
