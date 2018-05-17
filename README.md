# quagga-failover

In order to run a suffeciant ICMP network for different services we need to ensure the system stop advertising the loopback in case the local service fails on one of the hosts.

We use:
1. monit - for local services monitoring
2. quagga - virtual router
3. bgp-updater - simple failover script we wrote



monit is normally checking the local service and initiating "bgp-updater -state=down" if the service is down and "bgp-updater -state=up" when service is back up



