use strict;
#use warnings; FATAL => 'all';
while (<>) {
    my ($threadID, $requestType, $padding, $requesterID, $cmid, $timestamp, $ip, $packedCookies) =
        unpack("l!CCSa36La4a*", $_);
    print "threadID = $threadID,
        requestType = $requestType,
        padding = $padding,
        requesterID = $requesterID,
        cmid = $cmid,
        timestamp = $timestamp,
        ip = $ip,
        packedCookies = $packedCookies
    "
}
