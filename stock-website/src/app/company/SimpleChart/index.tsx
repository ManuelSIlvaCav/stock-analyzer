"use client";
import Script from "next/script";

export default function SimpleChart() {
  return (
    <div id="tradingview-wrapper" className="flex flex-col h-[50vh] p-4">
      <div className="text-center">Simple Chart</div>
      <div className="tradingview-widget-container__widget" />
      <Script
        id="tradingview-widget-advanced-chart"
        onLoad={() => {
          document
            ?.getElementById("tradingview-wrapper")
            ?.appendChild(
              document?.getElementById?.(
                "tradingview-widget-advanced-chart"
              ) as any
            );
        }}
        src="https://s3.tradingview.com/external-embedding/embed-widget-advanced-chart.js"
        type="text/javascript"
        strategy="lazyOnload"
        async={true}
      >
        {JSON.stringify({
          symbol: "NASDAQ:ULTA",
          interval: "D",
          timezone: "Etc/UTC",
          theme: "dark",
          style: "1",
          locale: "en",
          allow_symbol_change: false,
          calendar: false,
          support_host: "https://www.tradingview.com",
        })}
      </Script>
    </div>
  );
}
