To solve the billboard problem we're gonna use a divide and conquer approach.

In order to manage that we're gonna need a range within which the biggest fontSize possible is contained.

We're gonna use a simple heureustic as the starting point, we're gonna calculate the following ratio : max(W, H) / min(W, H).

With this ratio we'll calculate the starting font_size like this : min(W, H) / ratio, or max(W, H) / ratio if ratio > min(W, H) because we don't want to start with a fontsize of 0.

Once we have a starting point, we check whether or not we can fit the text on the billboard with the current fontSize.
If we can, we'll simply double the fontSize and check again until we can no longer fit the text on the billboard. In this case the last value will be the upper bound and the previous one the lower bound within which the optimal fontsize is.

If we can't, we'll half the fontSize and check again until we can fit the text on the billboard. The last value will be the lower bound and the preivous one the upper bound within which the optimal fontsize is.

The algorithm we're gonna use to find the range is the following :

```
if fits_on_billboard():
	do
		previous_size = font_size
		font_size = font_size * 2
	while fits_on_billboard():
	return [previous_size, font_size]
else:
	do
		previous_size =	font_size
		font_size = font_size / 2
	while not fits_on_billboard():
	return [font_size, previous_size]
fi

```

Once we have the range we can easily solve the problem:

```
range [min, max]
font_size = (min + max)/2
while min < font_size:
	if fits_on_billboard():
		min = font_size
	else: 
		max = font_size
	fi
	font_size = (min + max)/2
done
return min

```

At the end min will give us the right font_size
