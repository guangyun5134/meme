# Meme 代币发射平台

一个整合代币创建、曲线交易、行情展示和链上事件索引的全栈 Launchpad 项目。

## 项目特点

- 使用签名授权创建代币，并通过 bonding curve 提供初始买卖机制。
- 达到条件后将流动性迁移到 V2 AMM 交易池。
- 支持代币线性归属与领取，用于配置团队或其他参与方的释放计划。
- Go API 和 Indexer 分工处理业务请求、链上事件、成交记录及 K 线数据。
- 前端提供代币创建、买卖、行情图表和持币信息展示。

## 技术与目录

- `meme-launchpad/`：Solidity、Foundry 智能合约与测试。
- `meme-launchpad-back/`：Go、go-zero、PostgreSQL、Redis 服务。
- `meme-launchpad-fe/`：Next.js、React、TypeScript、Wagmi、Viem 前端。

克隆时使用 `git clone --recurse-submodules https://github.com/guangyun5134/meme.git`，以获取合约依赖。详细启动与配置见各模块 README。

项目包含教学代码或第三方依赖，保留原有来源说明与许可证。

RPC 与区块浏览器配置中的 `YOUR_...` 是占位符，运行或部署前请填写自己的配置。
