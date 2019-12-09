#!/usr/bin/env perl

use 5.30.0;
use strict;

my $input = <>;
chomp($input);

my @pixels                  = split //, $input;
my $num_of_pixels_per_layer = 25 * 6;
my $num_of_layers           = scalar @pixels / $num_of_pixels_per_layer;

my @img_pixels;
for ( my $l = $num_of_layers - 1 ; $l >= 0 ; --$l ) {
    for ( my $p = 0 ; $p < $num_of_pixels_per_layer ; ++$p ) {
        my $pixel = $pixels[ $l * $num_of_pixels_per_layer + $p ];

        if ( $pixel != 2 ) {
            $img_pixels[$p] = $pixel;
        }
    }
}

while (@img_pixels) {
    say join '', map { $_ ? '#' : ' ' } splice( @img_pixels, 0, 25 );
}
