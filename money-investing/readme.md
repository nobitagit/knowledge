# Money and Investing

## Compounding

Taking the interest gained on an investment and re-investing this portion into the same investment.

If you invert 10k in with 10% yearly return, here's the y.o.y. breakdown of the first 5 years.

| year | capital | interest |
-------|---------|-----------
| 1    | 11,000  | 1,000    |
| 2    | 12,000  | 1,000    |
| 3    | 13,000  | 1,000    |
| 4    | 14,000  | 1,000    |
| 5    | 15,000  | 1,000    |

With compounding:

| year | capital | interest |
-------|---------|-----------
| 1    | 11,000  | 1,000    |
| 2    | 12,100  | 1,100    |
| 3    | 13,310  | 1,210    |
| 4    | 14,641  | 1,331    |
| 5    | 16,105  | 1,464    |

So, we re-invest the interest each time to allow for compounding to take effect.
In 5 year the difference is between 15k and over 16k.

## Private vs Public corps

A **private corp** usually starts as a business with one person running and owning 100% of the business.
If Bob's business generates 100k a year, Bob gets to keep it all.

When Bob decides to grow the business he calculates he needs 2 millions. At this rate Bob will need 20+ years to grow enough capital. A possibility is to sell a percentage of the business, say 50%, to raise money and invest into growing.

If everything goes well, in 1 year the business grows so that it can generate 500k a year vs the original 100k.
Even if he indeed owns "less" of the company, the company now is worth more than originally, so Bob has actually increased his net worth even though he owns less shares.

While this works, it doesn't scale very well, because:

- finding an investor is time consuming
- Bob needs to vet the investor and trust them
- the investor needs to vet Bob and his business
- the paperwork takes time and is generally quite complex
- the value of the company is not easy to be determined and there's no arbitrage in place. Bob and the investor have to find **an agreement**

Going **public** streamlines this process. An investor buying shares:

- trusts that the revenue and financial info of the companies are truthful
- a public corp has to be audited every 3 months
- auditors certify that the books are correct

## IPO

An IPO is the process of bringing a company from private to public.
To do an IPO you need one or more **underwriters**, firms that specialise in bringing corps from private to public.
An underwriter:

- checks the books
- checks the future plans of the corp
- check the market prospects

They will then make an estimate like: "by selling 80% of the corp you can raise X amount of money"

NOTE: The number of shares of arbitrary and can be changed in time. Bob can decide to split the corp in 20000 shares or 7500000. Nothing really changes in terms of valuation.

So after the underwriters have done their analysis they strike a price per share and then contact big investors, big insurance firms and offer them part of the shares. This is the **primary market**.

Then, the shares are offered to the public for free trade. This is when the company actually goes public, by being traded in the **secondary market**.

## Stock exchange
It's the place where stocks can be traded. There are different stock exchanges, like NYSE, NASDAQ, London Stock Exchange.
Each ticker is listed and is trades on a specific SE. Apple trades on NASDAQ, the Bank of Montreal trades on the Toronto Stock Exchange. A company can also be interlisted, meaning it trades on multiple stock exchanges. In any case they will always be listed in 1 SE x country at most.

To trade on a SE though we can't just turn up ourselves, we need a **broker**.

As we've seen Bob has gone public. Let's assume:

- Stock IPO price: 28,75 $
- Shares outstanding: 7,000,000
- Exchange: NASDAQ

Even if Bob manages to sell the whole of the shares outstanding in the **primary market**, he will still have to go through the secondary market. This is a prerequisite and makes it possible for the shares to be traded in the first place.
Shares can actually be traded on secondary market before IPO, but it only happens if these stocks are not really worth it and if the primary market has decided that it's not a good investment.
In that case, it's practically always advisable to avoid buying.

### How are shares traded in practice

In the past brokers would trade in a physical place and would exchange and bargain by voice on the floors of the exchange. Now all exchanges are digital.
The system has an **order book**, which is similar to a long Excel sheet with 2 columns, **buyers** and **sellers**.

Now, John wants owns 700 shares and he wants to sell 100 of them. He calls his broker and has 2 options:

- **limit order**: John decides to sell (or buy) at a specific price. He is **not guaranteed** to be filled with this order (i.e. to completely sell the 100 shares)
- **market order**: John sells at whatever the price the market decides. He will sell at the best price available in that moment, and is guaranteed to be filled.

John wants to sell at 32$.
The broker fills in the order for John:

- **order type**: limit
- **quantity**: 100
- **side**: SELL (or BUY)
- **ticker**: SB (Bob's company's ticker)
- **price**: 32

In the past a physical person would go on the floor of NASDAQ and look for a buyer.
Now, the order enters the brokers system's servers and is sent to NASDAQ's servers.

The first thing that happens is here is that the system tries to match this order with an existing entry in the order book's buy side.
The order book is empty for now.

| Buyers | Sellers |
|--------|---------|
|        |         |
|        |         |
|        |         |
|        |         |

So the order is inserted in the sellers' side.

| Buyers | Sellers  |
|--------|----------|
|        |          |
|        |          |
|        | 32$ - 100|
|        |          |

Where 100 is the number of shares John puts on the market.
Marc also wants to sell SB shares. He wants to sell 750 at 35 dollars.
His order is placed in the book **after** John's one, because the book is **sorted by price**, with the best price on top.
Here's the state of the order book.

| Buyers | Sellers  |
|--------|----------|
|        |          |
|        |          |
|        | 32$ - 100|
|        |          |
|        | 35$ - 750|
|        |          |



### Quick notes

- Never buy stocks pre-IPO, if offered any
- Never buy in pre or after-market





























