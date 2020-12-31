# msgqcat

A trivial program to cat message queues.
Assumes the message does not contain newlines!

The perl packs them thusly:
my $packed = pack("l!CCSa36La4a*", $threadID, $requestType, 0, $requesterID, $cmid, $timestamp, $ip, $packedCookies);

See https://confluence.indexexchange.com/display/ADSG/ASCDB+-+Ad+Server+Core+Data+Base 

probably
while (<>) {
    my $threadID, $requestType, 0, $requesterID, $cmid, $timestamp, $ip, $packedCookies) = unpack("A10xA27xA7xA*", $_);
    print "$threadID, $requestType, 0, $requesterID, $cmid, $timestamp, $ip, $packedCookies"
}
