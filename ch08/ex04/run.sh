./reverb.out &
echo One | ./netcat.out &
echo Two | ./netcat.out &
wait
