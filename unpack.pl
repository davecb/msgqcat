use strict;
#use warnings; FATAL => 'all';
while (<>) {
    my ($threadID, $requestType, $junk, $requesterID, $cmid, $timestamp, $ip, $packedCookies) =
        unpack("l!CCSa36La4a*", $_);
    print "$threadID, $requestType, $junk, $requesterID, $cmid, $timestamp, $ip, $packedCookies)"
}
