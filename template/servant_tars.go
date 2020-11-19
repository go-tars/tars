package template

var (
	ServantTars = `module {{.App}}
{
	interface {{.Servant}}
	{
	    int Add(int a,int b,out int c); // Some example function
	    int Sub(int a,int b,out int c); // Some example function
	};
};
`
)