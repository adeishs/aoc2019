#!/usr/bin/env perl

use 5.28.0;
use List::Util qw(sum);
use Data::Dumper::Concise;

use constant COORDS => [qw(x y z)];

sub parse_input_line {
    my ($line) = @_;

    state $re_str =
      '<' . join( ', ', ( map { $_ . '=(-?\d+)' } COORDS->@* ) ) . '>';
    my (@positions) = $line =~ qr/$re_str/;

    return { map { $_ => shift @positions } COORDS->@* };
}

my @lines = <STDIN>;
my @moons;
push @moons, parse_input_line($_) for @lines;
for my $m (0 .. $#moons) {
    $moons[$m]->{'v' . $_} = 0 for COORDS->@*;
}

for (1 .. 1000) {
    my @tmp_moons = @moons;
    for my $i (0 .. $#moons) {
        for my $j ( grep { $_ != $i } 0 .. $#moons) {
            for my $coord (COORDS->@*) {
                $tmp_moons[$i]->{"v$coord"} += $moons[$j]->{$coord} <=> $moons[$i]->{$coord};
            }
        }
    }
    for my $i (0 .. $#moons) {
        for my $coord (COORDS->@*) {
            $moons[$i]->{"v$coord"} = $tmp_moons[$i]->{"v$coord"};
            $moons[$i]->{$coord} += $moons[$i]->{"v$coord"};
        }
    }
    @moons = @tmp_moons;
}

my $tot_e = 0;
for my $m (@moons) {
    my %e = 0;
    for my $c (COORDS->@*) {
        for ('', 'v') {
            $e{$_} += abs($m->{"$_$c"});
        }
    }
    $tot_e += $e{''} * $e{v};
}

say $tot_e;
