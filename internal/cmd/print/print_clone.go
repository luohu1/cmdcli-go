package print

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

type CloneOptions struct {
	token   string
	baseURL string
}

func newCmdPrintClone() *cobra.Command {
	opts := &CloneOptions{}

	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Print clone shell",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gitlab.NewClient(opts.token, gitlab.WithBaseURL(opts.baseURL))
			if err != nil {
				fmt.Printf("Failed to create client: %v", err)
			}

			opt := &gitlab.ListProjectsOptions{
				ListOptions: gitlab.ListOptions{
					Page:    1,
					PerPage: 110,
				},
				Simple: gitlab.Bool(true),
			}

			for {
				projects, resp, err := client.Projects.ListProjects(opt)
				if err != nil {
					panic(err)
				}

				for _, p := range projects {
					fmt.Printf("git clone %s %s\n", p.HTTPURLToRepo, p.PathWithNamespace)
				}

				if resp.CurrentPage >= resp.TotalPages {
					break
				}

				opt.Page = resp.NextPage
			}
		},
	}

	cmd.Flags().StringVarP(&opts.token, "token", "t", "", "Gitlab auth token")
	cmd.Flags().StringVarP(&opts.baseURL, "base-url", "b", "", "Gitlab base url")

	return cmd
}
