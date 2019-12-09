#!/usr/bin/env perl

use 5.30.0;
use strict;

my $input = <>;
chomp($input);

my @pixels                  = split //, $input;
my $num_of_pixels_per_layer = 25 * 6;
my $num_of_layers           = scalar @pixels / $num_of_pixels_per_layer;

my %freq;
for ( my $l = 0 ; $l < $num_of_layers ; ++$l ) {
    for ( my $p = 0 ; $p < $num_of_pixels_per_layer ; ++$p ) {
        my $pixel = $pixels[ $l * $num_of_pixels_per_layer + $p ];

        $freq{$l}->{$pixel}++;
    }
}

my $min_cnt = scalar @pixels;
my $min_cnt_layer;
for ( my $l = 0 ; $l < $num_of_layers ; ++$l ) {
    if ( $freq{$l}->{0} < $min_cnt ) {
        $min_cnt = $freq{$l}->{0};
        $min_cnt_layer = $l;
    }
}
say ($freq{$min_cnt_layer}->{1} * $freq{$min_cnt_layer}->{2});
