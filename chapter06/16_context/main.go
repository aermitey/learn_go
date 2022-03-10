package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//withCancel()
	//withTimeOut()
	//withValue()
	withDeadline()
}

func withCancel() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，买材料")
	go buyFlour(ctx)
	go buyOil(ctx)
	go buyEgg(ctx)
	time.Sleep(1 * time.Second)
	fmt.Println("没电了，不买了")
	cancel() //当调用cancel后，所有由此上下文衍生出来的context都会cancel
	time.Sleep(10 * time.Second)
}

func buyFlour(ctx context.Context) {
	fmt.Println("去买面")
	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买面")
		return
	default:
	}
	fmt.Println("买面")
}

func buyOil(ctx context.Context) {
	fmt.Println("去买油")
	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买油")
		return
	default:
	}
	fmt.Println("买油")
}

func buyEgg(ctx context.Context) {
	ctx1, _ := context.WithCancel(ctx)
	//defer cancel()//无论是否调用衍生出来的ctx的cancel，done返回的cancel都会关闭
	fmt.Println("去买蛋")
	//time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋")
		return
	default:
	}
	fmt.Println("买蛋")
	buySEgg(ctx1)
	buyBEgg(ctx1)
	time.Sleep(3 * time.Second)
}

func buySEgg(ctx context.Context) {
	fmt.Println("去买鸽子蛋")
	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买鸽子蛋")
		return
	default:
	}
	fmt.Println("买鸽子蛋")
}

func buyBEgg(ctx context.Context) {
	fmt.Println("去买鸵鸟蛋")
	time.Sleep(2 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买鸵鸟蛋")
		return
	default:
	}
	fmt.Println("买鸵鸟蛋")
}

func withTimeOut() {
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	fmt.Println("开始部署望远镜")
	go distributeMainFrame(ctx)
	go distributeMainBody(ctx)
	go distributeCover(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("任务超时没有完成")
	}
	time.Sleep(10 * time.Second)
}

func distributeMainFrame(ctx context.Context) {
	time.Sleep(1 * time.Second)
	fmt.Println("distributeMainFrame")
	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainFrame")
		return
	default:
	}
	//time.Sleep(10 * time.Second)

}

func distributeMainBody(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainBody")
		return
	default:
	}
	fmt.Println("distributeMainBody")
}

func distributeCover(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeCover")
		return
	default:
	}
	fmt.Println("distributeCover")
}

func withValue() {
	ctx := context.WithValue(context.TODO(), "1", "钱包")
	goToPapa(ctx)
	fmt.Println("withValue:1", ctx.Value("1"))
	fmt.Println("withValue:2", ctx.Value("2"))
	fmt.Println("withValue:3", ctx.Value("3"))
	fmt.Println("withValue:4", ctx.Value("4"))
}

func goToPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, "2", "充电宝")
	goToMama(ctx)
	fmt.Println("goToPapa:1", ctx.Value("1"))
	fmt.Println("goToPapa:2", ctx.Value("2"))
	fmt.Println("goToPapa:3", ctx.Value("3"))
	fmt.Println("goToPapa:4", ctx.Value("4"))
}

func goToMama(ctx context.Context) {
	ctx = context.WithValue(ctx, "3", "小夹克")
	goToGrandma(ctx)
	fmt.Println("goToMama:1", ctx.Value("1"))
	fmt.Println("goToMama:2", ctx.Value("2"))
	fmt.Println("goToMama:3", ctx.Value("3"))
	fmt.Println("goToMama:4", ctx.Value("4"))
}

func goToGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "大苹果")
	goToPark(ctx)
	fmt.Println("goToGrandma:1", ctx.Value("1"))
	fmt.Println("goToGrandma:2", ctx.Value("2"))
	fmt.Println("goToGrandma:3", ctx.Value("3"))
	fmt.Println("goToGrandma:4", ctx.Value("4"))
}

func goToPark(ctx context.Context) {
	fmt.Println("goToPark:1", ctx.Value("1"))
	fmt.Println("goToPark:2", ctx.Value("2"))
	fmt.Println("goToPark:3", ctx.Value("3"))
	fmt.Println("goToPark:4", ctx.Value("4"))
}

func withDeadline() {
	ctx, _ := context.WithDeadline(context.TODO(), time.Now().Add(2*time.Second))
	go watchTV(ctx)
	go watchMobile(ctx)
	go playGame(ctx)
	time.Sleep(3 * time.Second)
}

func watchTV(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电视")
			return
		default:
		}
		fmt.Println("看电视")
		time.Sleep(300 * time.Millisecond)
	}
}

func watchMobile(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关手机")
			return
		default:
		}
		fmt.Println("看手机")
		time.Sleep(300 * time.Millisecond)
	}
}

func playGame(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关游戏机")
			return
		default:
		}
		fmt.Println("玩游戏")
		time.Sleep(300 * time.Millisecond)
	}
}
