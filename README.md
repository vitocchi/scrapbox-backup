# scrapbox-backup
AWS Lambdaで運用するLambda関数をGolangで作成した。
一日一回、scrapboxのプロジェクトのバックアップをAWS S3に保存する

KMSでGitHubTokenを暗号化しているので、暗号化したTOKENをそのままgithubに上げれちゃう。便利。

1. GitHubでAccess Tokenを取得

2. aws cliでkmsのAPIを叩いてAccess Tokenを暗号化する

3. ソースコード内でkmsのSDKを使って暗号化したTOKENを復号化する

4. Lambda関数にKMSを利用するロールをもたせることでLambdaは復号化APIを叩くことができる

5. KMSを利用するロールを持っていなければAccess Tokenを復号化出来ない。

今回はGithubを使わなかったが、scrapboxの認証クッキーについて同様の操作を行っている
