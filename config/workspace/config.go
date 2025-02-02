package workspace

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("singlestoredb_workspace", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "singlestoredb"
		r.ShortGroup = "singlestore"
		r.Kind 		 = "Workspace"
		r.References["workspace_group_id"] = config.Reference{
			TerraformName: "singlestoredb_workspace_group",
		}
	})
}
