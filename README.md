# 1

## 服务创建步骤

- 创建/service/product/model 文件夹并创建product.sql 文件

```sql
CREATE TABLE `product` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '产品名称',
    `desc` varchar(255)  NOT NULL DEFAULT '' COMMENT '产品描述',
        `stock` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '产品库存',
    `amount` int(10) unsigned NOT NULL DEFAULT '0'  COMMENT '产品金额',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '产品状态',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
```

- cd到service 目录

```sh
# model
goctl model mysql ddl -src ./model/product.sql -dir ./model -c
# api
goctl api go -api ./api/product.api -dir ./api
# rpc
goctl rpc protoc product.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
```

- 创建api /service/product/model product.api

```go
type (
    // 产品创建
    CreateRequest {
        Name   string `json:"name"`
        Desc   string `json:"desc"`
        Stock  int64  `json:"stock"`
        Amount int64  `json:"amount"`
        Status int64  `json:"status"`
    }
    CreateResponse {
        Id int64 `json:"id"`
    }
    // 产品创建

    // 产品修改
    UpdateRequest {
        Id     int64  `json:"id"`
        Name   string `json:"name,optional"`
        Desc   string `json:"desc,optional"`
        Stock  int64  `json:"stock"`
        Amount int64  `json:"amount,optional"`
        Status int64  `json:"status,optional"`
    }
    UpdateResponse {
    }
    // 产品修改

    // 产品删除
    RemoveRequest {
        Id int64 `json:"id"`
    }
    RemoveResponse {
    }
    // 产品删除

    // 产品详情
    DetailRequest {
        Id int64 `json:"id"`
    }
    DetailResponse {
        Id     int64  `json:"id"`
        Name   string `json:"name"`
        Desc   string `json:"desc"`
        Stock  int64  `json:"stock"`
        Amount int64  `json:"amount"`
        Status int64  `json:"status"`
    }
    // 产品详情
)

@server(
    jwt: Auth
)
service Product {
    @handler Create
    post /api/product/create(CreateRequest) returns (CreateResponse)

    @handler Update
    post /api/product/update(UpdateRequest) returns (UpdateResponse)

    @handler Remove
    post /api/product/remove(RemoveRequest) returns (RemoveResponse)

    @handler Detail
    post /api/product/detail(DetailRequest) returns (DetailResponse)
}
```

- 创建rpc /service/product/pb/product.proto

```proto
syntax = "proto3";

package productclient;

option go_package = "gzero-user/service/product";

// 产品创建
message CreateRequest {
    string Name = 1;
    string Desc = 2;
    int64 Stock = 3;
    int64 Amount = 4;
    int64 Status = 5;
}
message CreateResponse {
    int64 id = 1;
}
// 产品创建

// 产品修改
message UpdateRequest {
    int64 id = 1;
    string Name = 2;
    string Desc = 3;
    int64 Stock = 4;
    int64 Amount = 5;
    int64 Status = 6;
}
message UpdateResponse {
}
// 产品修改

// 产品删除
message RemoveRequest {
    int64 id = 1;
}
message RemoveResponse {
}
// 产品删除

// 产品详情
message DetailRequest {
    int64 id = 1;
}
message DetailResponse {
    int64 id = 1;
    string Name = 2;
    string Desc = 3;
    int64 Stock = 4;
    int64 Amount = 5;
    int64 Status = 6;
}
// 产品详情

service Product {
    rpc Create(CreateRequest) returns(CreateResponse);
    rpc Update(UpdateRequest) returns(UpdateResponse);
    rpc Remove(RemoveRequest) returns(RemoveResponse);
    rpc Detail(DetailRequest) returns(DetailResponse);
}
```
