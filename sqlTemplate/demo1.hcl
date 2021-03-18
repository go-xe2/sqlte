// 查询家订单列表
dboName = "demo1"

select {
  queryRows = <<SQL
      select * from mny_order_master where order_id in(
        select order_id from mny_buyer_order where buyer_id = :buyer_id
        {{if .page -}}
           LIMIT {{.page.Offset}},{{.page.PageSize -}}
        {{else -}}
           LIMIT 0, 30
        {{end -}}
      )
      {{if ge .status 0 -}}
        and status = :status
      {{end -}}
      {{if ge .start_date 0 -}}
        and order_date >= :start_date
      {{end -}}
      {{if ge .end_date 0 -}}
        and order_date <= :end_date
      {{end -}}
    SQL

  queryRowCount = <<SQL
    select count(1) from mny_order_master where order_id in(
      select order_id from mny_buyer_order where buyer_id = :buyer_id
      {{ if .page }}
        LIMIT {{.page.Offset}},{{.page.PageSize}}
      {{ else }}
        LIMIT 0, 30
      {{ end }}
    )
    {{ if .status ge 0 }}
      and status = :status
    {{ end }}
    {{ if .start_date ge 0 }}
      and order_date >= :start_date
    {{ end }}
    {{ if .end_date ge 0 }}
      and order_date <= :end_date
    {{ end }}
  SQL
}

execute {
  deleteBuyOrder = <<SQL
    delete from mny_buyer_order where order_id = ''
  SQL
}
