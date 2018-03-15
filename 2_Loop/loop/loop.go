package loop

// basic loop
func Basic(sum int) int{
	for i:=0; i<10; i++{
		sum += sum;
	}

	return sum
}

// loop without initialization and post increments
func Optional(sum int) int{
	for ; sum<100; {
		sum += sum;
	}

	return sum
}

// forever loop
func Forever(sum int) int{
	for{
		sum += sum
		if sum >= 40{
			break
		}
	}

	return sum
}