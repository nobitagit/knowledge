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
His order is placed in the book **after** John's one, because the book SELL side is **sorted by price**, with the best price on top.
Here's the state of the order book.

| Buyers | Sellers  |
|--------|----------|
|        |          |
|        |          |
|        | 32$ - 100|
|        |          |
|        | 35$ - 750|
|        |          |

Tom wants to buy, but only at 29$. His order will enter the book:

| Buyers | Sellers  |
|--------|----------|
|        |          |
|70 - 29$|          |
|        | 32$ - 100|
|        |          |
|        | 35$ - 750|
|        |          |

From this state, we can derive the **bid and ask**.
The bid is the highest price a person is willing to pay for to buy a single share.
The ask is the lowest price a person is willing to pay for to sell a single share.
In the case above:

- bid is 29
- ask is 32

The **last trade** is the price of the single share in the latest transaction. There hasn't been transactions so far, so last trade price is available.
Another order comes in:


| Buyers | Sellers  |
|--------|----------|
|30 - 31$|          |
|70 - 29$|          |
|        | 32$ - 100|
|        |          |
|        | 35$ - 750|
|        |          |

As we see, the BUY side of the book is **sorted by price too**, but with the highest price on top.
Bid and ask changes:

- bid is 31
- ask is 32

As we see, B&A is always the 2 values on top.

Carol wants to buy now. She sends in a **market order**, so she will buy no matter what, paying the best price in that moment.

- **order type**: market
- **quantity**: 50
- **side**: BUY
- **ticker**: SB (Bob's company's ticker)
- **price**: -- (no price for market orders)

This will change the order book:

| Buyers | Sellers  |
|--------|----------|
|30 - 31$|          |
|70 - 29$|          |
|        | 32$ - 50 |
|        |          |
|        | 35$ - 750|
|        |          |

50 shares have been deducted from the topmost entry in the sell side.
Carol is filled, her order is complete and was executed at 32$ per shares.
John is **partially filled**, his order to sell has been partially executed, he still has 50 shares left to sell.

- bid is 31$
- ask is 32$
- last trade is 32$

A new order comes in, Paul wants to buy:

- **order type**: limit
- **quantity**: 140
- **side**: BUY
- **ticker**: SB (Bob's company's ticker)
- **price**: 33

The order books assigns the remaining 50 shares from John @ 32$. Notice how Paul is willing to pay 33, but still pays only 32 for the first 50 shares.
John is completely filled now.
Paul though still wants 90 more shares, so his order is added to the order book with a request of buying 90 (140-50) shares at 33.

| Buyers | Sellers  |
|--------|----------|
|90 - 33$|          |
|30 - 31$|          |
|70 - 29$|          |
|        |          |
|        | 35$ - 750|
|        |          |

- bid is 33$
- ask is 35$
- last trade is 32

## Market cap

It's the total number of outstanding shares * the value of each share.
The MC of each company goes up and down every day based on fluctuations.

**Blue chips** are companies with a big market cap (10 to 200 billions).
The lower the cap, the higher the risk.

> These large companies have usually been around for a long time, and they are major players in well-established industries. Investing in large-cap companies does not necessarily bring in huge returns in a short period of time, but over the long run, these companies generally reward investors with a consistent increase in share value and dividend payments. An example of a large-cap company is International Business Machines (IBM),1 Johnson & Johnson (JNJ),2 or Microsoft (MSFT).3
>
> Mid-cap companies generally have a market capitalization of between $2 billion and $10 billion. Mid-cap companies are established companies that operate in an industry expected to experience rapid growth. Mid-cap companies are in the process of expanding. They carry an inherently higher risk than large-cap companies because they are not as established, but they are attractive for their growth potential. An example of a mid-cap company is Eagle Materials (EXP).4

Every quarter companies have to release **financial statements**. They release:

- Balance sheet
- Income statement: the document that is behind the earnings
- Cash flow statement: where does the cash flow come from? From core operations or one off movements (like investing or selling assets?). Important to understand if the core business is really generating money

These are verified by auditors.

## Earnings

[link](https://www.udemy.com/course/complete-investing-course-stocks-etfs-index-mutual-funds/learn/lecture/22963892?start=15#questions)

- E or Earnings: the profit made by a company in a period (a quarter)
- EPS or Earnings per share: E divided by the n. of shares outstanding.
- PE ration or Price to earning: the ratio of the stock share price divided by the stock EPS. It tells you how much you are paying for each dollar a company makes. It flags if a company is overvalued or undervalued.

If company A earns 3,000,000 and company B earns 9,000,000 from which company of the 2 would you buy 1 share?
You can't tell with earnings alone because you would need to know the amount of shares outstanding.

If A has 5,000 shares but B has 90,000 shares the value of 1 share differs dramatically in favour of A.
EPS will then be calculated like so:

- A: 3000000/5000 = 600
- B: 9000000/90000 = 100

So, here we see that A is actually a more desirable share to have, because it earns more.
BUT still PE ratio is a better indicator of the worth of a stock because 2 shares with same EPS (say 30), can cost a very different amount.
So PE will tell you that. A lower value is better.

## Book value

It's the value value of the business according to its books. If a company had to sell all its assets and pay all its liabilities, the final amount would be its book value.
**BVPS** it's the book value divided by the n. of outstanding shares.
**Price to book** or **P/B** is calculated by dividing the stock price by its BVPS. It tells us **how much we're paying for each dollar of book value**.

Here are a few PB values:

- On 11.52
- Tesla 37.03
- Walmart 4.77
- Amazon 14.48
- Coca Cola 5.64
- GE 2.84

## Growth vs Value stocks

Growth companies: companies that are not generating as much income relative to the share price but are expected to have high earnings growth in the future. They usually have a high PE ratio and do not give dividends. They have high PB. An example, Tesla.

Value stocks: more established companies that have good income relative to the share price. Generally low PE ratio and pay out dividends. They have lower PB. An example, Walmart, Coca Cola.

The difference in PB is due - in general - to the fact that value companies have more assets while growth companies are mostly values because of their **potential**, so there's more risk.

## Dividends

A dividend is a distribution of profits by a corp to its shareholders. Any profit not distributed is re-invested in the company itself.
The corp announces when they will pay dividends on a day called **declaration date**. Most corps pay quarterly dividends.

### Quick notes

- Never buy stocks pre-IPO, if offered any
- Never buy in pre or after-market





























