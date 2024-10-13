export namespace main {
	
	export class Config {
	    github_token: string;
	    table_name: string;
	    owner: string;
	    repo: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.github_token = source["github_token"];
	        this.table_name = source["table_name"];
	        this.owner = source["owner"];
	        this.repo = source["repo"];
	    }
	}
	export class Release {
	    tag_name: string;
	    name: string;
	    created_at: string;
	    html_url: string;
	    Current: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Release(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tag_name = source["tag_name"];
	        this.name = source["name"];
	        this.created_at = source["created_at"];
	        this.html_url = source["html_url"];
	        this.Current = source["Current"];
	    }
	}
	export class VersionedData {
	    Releases: Release[];
	    Version: string;
	
	    static createFrom(source: any = {}) {
	        return new VersionedData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Releases = this.convertValues(source["Releases"], Release);
	        this.Version = source["Version"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

