package docker

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	// Docker CLI bilan bog‘lanish
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	// Docker versiyasini tekshirish
	info, err := cli.ServerVersion(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🔹 Docker versiyasi:", info.Version)

	// Yangi container yaratish
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",                                    // Alpine Linux image'ini ishlatamiz (eng yengil OS)
		Cmd:   strslice.StrSlice{"echo", "Hello, Docker!"}, //Container ichida buyruq: `echo "Hello, Docker!"`
	}, nil, nil, nil, "my_alpine_container")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Container yaratildi:", resp.ID)

	// Containerni ishga tushirish
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("🚀 Container ishga tushdi:", resp.ID)

	// Container jarayonlarini olish
	logs, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("📜 Container loglari:")
	defer logs.Close()
	logReader := make([]byte, 1024)
	n, _ := logs.Read(logReader)
	fmt.Println(string(logReader[:n]))

	// Containerni to‘xtatish va o‘chirish
	if err := cli.ContainerStop(ctx, resp.ID, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Container to‘xtatildi:", resp.ID)

	if err := cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Container o‘chirildi:", resp.ID)
}
